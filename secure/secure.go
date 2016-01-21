package secure

import (
	"os"
    "crypto/rsa"
    "crypto/rand"
    "crypto/x509"
    //"golang.org/x/crypto/ssh"
    "encoding/pem"
    "io/ioutil"
)

/*
Functions in this package help with generating and loading keys for use by
the overlord, master, and peon.  They are used to secure communication
*/

func SetupKeys() error {
    mkDirs("/etc/golem/pki/master")
    err := MakeSSHKeyPair("/etc/golem/pki/master/master_pub.pub", "/etc/golem/pki/master/master_key.pem")
    return err 
}

// MakeSSHKeyPair make a pair of public and private keys for SSH access.
// Public key is encoded in the format for inclusion in an OpenSSH authorized_keys file.
// Private Key generated is PEM encoded
func MakeSSHKeyPair(pubKeyPath, privateKeyPath string) error {
    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        return err
    }

    // generate and write private key as PEM
    privateKeyFile, err := os.Create(privateKeyPath)
    defer privateKeyFile.Close()
    if err != nil {
        return err
    }
    privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
    if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
        return err
    }

    // generate and write public key
    publicKeyPEM, err := x509.MarshalPKIXPublicKey(privateKey.PublicKey)
    //pub, err := ssh.NewPublicKey(&privateKey.PublicKey)
    if err != nil {
        return err
    }
    return ioutil.WriteFile(pubKeyPath, publicKeyPEM, 0655)
}

// Exists reports whether the named file or directory exists.
func mkDirs(names ...string) {
    for _, name := range names {
        if _, err := os.Stat(name); err != nil {
            if os.IsNotExist(err) {
                os.MkdirAll(name, 0400)
            }
        }
    }
}

