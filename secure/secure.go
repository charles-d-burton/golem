package secure

import (
	"log"
	"os"
	"io/ioutil"
)

const (
	masterKeyDir     = "/etc/golem/pki/master"
	MasterPrivateKey = "/etc/golem/pki/master/master_key"
	MasterPubCert     = "/etc/golem/pki/master/master_pub"

	MasterAcceptDir  = "/etc/golem/pki/master/peons"
	MasterPendingDir = "/etc/golem/pki/master/pending"

	peonKeyDir     = "/etc/golem/pki/peon"
	PeonPrivateKey = "/etc/golem/pki/peon/peon_key"
	PeonPubCert     = "/etc/golem/pki/peon/peon_pub"
)

type keys struct {
	Pub  string
	Priv string
}

/*
Functions in this package help with generating and loading keys for use by
the overlord, master, and peon.  They are used to secure communication by encrypting
the data that's sent.
*/
func SetupKeys(mode string) error {
	if mode == "master" {
		mkDirs(masterKeyDir)
		mkDirs(MasterAcceptDir)
		if !checkKeys(MasterPrivateKey) || !checkKeys(MasterPubCert) {
			log.Println("Master Keys Not Found, Generating new keys ...")
			cert, key, err := GenerateMemCert()
			err = ioutil.WriteFile(MasterPrivateKey, key, 0640)
			err = ioutil.WriteFile(MasterPubCert, cert, 0640)
			if err != nil {
				log.Println(err.Error())
				return err
			}
			
		} else {
			log.Println("Keys Found")
		}
	} else if mode == "peon" {
		mkDirs(peonKeyDir)
		if !checkKeys(PeonPrivateKey) || !checkKeys(PeonPubCert) {
			log.Println("Peon Keys Not Found, Generating new keys ...")
			cert, key, err := GenerateMemCert()
			err = ioutil.WriteFile(PeonPrivateKey, key, 0640)
			err = ioutil.WriteFile(PeonPubCert, cert, 0640)
			if err != nil {
				log.Println(err.Error())
				return err
			}
		} else {
			log.Println("Keys Found")
		}
	}
	return nil
}

/*
func SetupKeys(mode string) error {
	if mode == "master" {
		mkDirs(masterKeyDir)
		mkDirs(MasterAcceptDir)
		if !checkKeys(MasterPrivateKey) || !checkKeys(MasterPubKey) {
			log.Println("Master Keys Not Found, Generating new keys ...")
			masterPub, masterKey, _ := box.GenerateKey(rand.Reader)
			err := ioutil.WriteFile(MasterPubKey, masterPub[:], 0640)
			err = ioutil.WriteFile(MasterPrivateKey, masterKey[:], 0640)
			if err != nil {
				log.Println(err.Error())
				return err
			}
			//err := MakeSSHKeyPair("/etc/golem/pki/master/master_pub.pem", "/etc/golem/pki/master/master_key.pem")
			//return err
		}
		log.Println("Keys Found")
	} else if mode == "peon" {
		mkDirs(peonKeyDir)
		if !checkKeys(PeonPrivateKey) || !checkKeys(PeonPubKey) {
			log.Println("Peon Keys Not Found, Generating new keys ...")
			peonPub, peonKey, _ := box.GenerateKey(rand.Reader)
			err := ioutil.WriteFile(PeonPubKey, peonPub[:], 0640)
			err = ioutil.WriteFile(PeonPrivateKey, peonKey[:], 0640)
			if err != nil {
				log.Println(err.Error())
				return err
			}
		}
	}
	return nil
}*/

// MakeSSHKeyPair make a pair of public and private keys for SSH access.
// Public key is encoded in the format for inclusion in an OpenSSH authorized_keys file.
// Private Key generated is PEM encoded
/*func MakeSSHKeyPair(pubKeyPath, privateKeyPath string) error {
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
}*/

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
