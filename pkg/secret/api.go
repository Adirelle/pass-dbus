package secret

import (
	"time"

	"github.com/godbus/dbus/v5"
)

type (
	Fields map[string]string

	CollectionMethods interface {
		Delete() (err *dbus.Error)
		SearchItems(fields Fields) (results []dbus.ObjectPath, err *dbus.Error)
		CreateItem(fields Fields, secret *Secret, label string, replace bool) (item dbus.ObjectPath, err *dbus.Error)
	}

	CollectionProperties struct {
		Items    []dbus.ObjectPath
		Private  string
		Label    string
		Locked   string
		Created  time.Time
		Modified time.Time
	}

	CreatedItem struct {
		Item dbus.ObjectPath
	}

	DeletedItem struct {
		Item dbus.ObjectPath
	}

	ItemMethods interface {
		Delete() (err *dbus.Error)
	}

	ItemProperties struct {
		Locked     bool
		Attributes Fields
		Label      string
		Secret     *Secret
		Created    time.Time
		Modified   time.Time
	}

	ItemChanged struct{}

	ServiceProperties struct {
		Collections       []dbus.ObjectPath
		DefaultCollection dbus.ObjectPath
	}

	ServiceInterface interface {
		OpenSession() (session *dbus.ObjectPath, err *dbus.Error)
		CreateCollection(label string, private bool) (err *dbus.Error)
		LockService() (err *dbus.Error)
		SearchCollections(fields Fields) (results []dbus.ObjectPath, locked []dbus.ObjectPath, err *dbus.Error)
		RetrieveSecrets(items []dbus.ObjectPath) (secrets []*Secret, err *dbus.Error)
	}

	CollectionCreated struct {
		Collection dbus.ObjectPath
	}

	CollectionDeleted struct {
		Collection dbus.ObjectPath
	}

	SessionInterface interface {
		Close(err *dbus.Error)
		Negotiate(algorithm string) (output *dbus.Variant, err *dbus.Error)
		BeginAuthenticate(objects []dbus.ObjectPath, windowId string) (err *dbus.Error)
		CompleteAuthenticate(objects []dbus.ObjectPath, authenticated []dbus.ObjectPath) (err *dbus.Error)
	}

	Authenticated struct {
		Object  dbus.ObjectPath
		Success bool
	}

	Secret struct {
		Algorithm  string
		Parameters []byte
		Value      []byte
	}
)

var (
	ErrAlreadyExists = dbus.NewError("org.freedesktop.Secrets.Error.AlreadyExists", nil)
	ErrIsLocked      = dbus.NewError("org.freedesktop.Secrets.Error.IsLocked", nil)
	ErrNotSupported  = dbus.NewError("org.freedesktop.Secrets.Error.NotSupported", nil)
)
