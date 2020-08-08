/*
Copyright AppsCode Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package verifier

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func GetClusterUID(client corev1.NamespaceInterface) (string, error) {
	ns, err := client.Get(context.TODO(), metav1.NamespaceSystem, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return string(ns.UID), nil
}

func VerifyLicense(product, clusterUID string, caCert []byte, license []byte) error {
	block, _ := pem.Decode(license)
	if block == nil {
		// This probably is a JWT token, should be check for that when ready
		return errors.New("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return errors.Wrap(err, "failed to parse certificate")
	}

	// First, create the set of root certificates. For this example we only
	// have one. It's also possible to omit this in order to use the
	// default root set of the current operating system.
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(caCert)
	if !ok {
		return errors.New("failed to parse root certificate")
	}

	opts := x509.VerifyOptions{
		DNSName: clusterUID,
		Roots:   roots,
		KeyUsages: []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth,
		},
	}
	if _, err := cert.Verify(opts); err != nil {
		return errors.Wrap(err, "failed to verify certificate")
	}
	if !sets.NewString(cert.Subject.Organization...).Has(product) {
		return fmt.Errorf("license was not issued for %s", product)
	}
	return nil
}
