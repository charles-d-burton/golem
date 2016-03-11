package secure

import (
	"os"
    "crypto/rsa"
    "crypto/rand"
    "crypto/x509"
    "encoding/pem"
    "io/ioutil"
    "log"
)

type keys struct {
    Pub string
    Priv string
}

/*
Functions in this package help with generating and loading keys for use by
the overlord, master, and peon.  They are used to secure communication
*/

func SetupKeys() error {
    mkDirs("/etc/golem/pki/master")
    if !checkKeys("/etc/golem/pki/master/master_pub.pem") || !checkKeys("/etc/golem/pki/master/master_key.pem") {
        log.Println("Keys Not Found, Generating new keys ...")
        err := MakeSSHKeyPair("/etc/golem/pki/master/master_pub.pem", "/etc/golem/pki/master/master_key.pem")
        return err
    }
    log.Println("Keys Found")
    return nil 
}

// MakeSSHKeyPair make a pair of public and private keys for SSH access.
// Public key is encoded in the format for inclusion in an OpenSSH authorized_keys file.
// Private Key generated is PEM encoded
func MakeSSHKeyPair(pubKeyPath, privateKeyPath string) error {
    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        return err
    }
    err = privateKey.Validate();
	if err != nil {
		log.Println("Validation failed.", err);
	}

	// Get der format. privDer []byte
	privDer := x509.MarshalPKCS1PrivateKey(privateKey);

	// pem.Block
	// blk pem.Block
	privBlk := pem.Block {
	Type: "RSA PRIVATE KEY",
	Headers: nil,
	Bytes: privDer,
	};
    

	// Resultant private key in PEM format.
    err = ioutil.WriteFile(privateKeyPath, pem.EncodeToMemory(&privBlk), 0640);
    if err != nil {
        return err
    }
	// Public Key generation

	pub := privateKey.PublicKey;
	pubDer, err := x509.MarshalPKIXPublicKey(&pub);
	if err != nil {
		log.Println("Failed to get der format for PublicKey.", err);
		return err;
	}

	pubBlk := pem.Block {
	Type: "PUBLIC KEY",
	Headers: nil,
	Bytes: pubDer,
	}
	//pub_pem := string(pem.EncodeToMemory(&pubBlk));
	//log.Printf(pub_pem);
    return ioutil.WriteFile(pubKeyPath, pem.EncodeToMemory(&pubBlk), 0644)
}

// mkDirs reports whether the named file or directory exists.
func mkDirs(names ...string) {
    for _, name := range names {
        if _, err := os.Stat(name); err != nil {
            if os.IsNotExist(err) {
                os.MkdirAll(name, 0400)
            }
        }
    }
}

//Check to see if the keys have already been generated
func checkKeys(name string) bool {
    if _, err := os.Stat(name); err == nil {
        log.Println(name, "exist!")
        return true
    }
    
    return false
}

