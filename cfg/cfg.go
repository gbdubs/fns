package cfg

import (
	"encoding/json"
	"io/ioutil"
)

type Config interface {
	cfg() *configData
	fileName() string
}

type configData struct {
	origin string
	root   string
}

const originLocal string = "local"
const originNew string = "new"
const originDerivedLocal string = "derived-local"
const originDerivedNew string = "derived-new"

func (c configData) Root() string {
	return c.root
}

func Root(c Config) string {
	return c.cfg().Root()
}
func Path(c Config) string {
	return c.cfg().Root() + "/" + c.fileName()
}

func (c *configData) existsOnDisk() bool {
	return c.origin == originLocal || c.origin == originDerivedLocal
}

func ExistsOnDisk(c Config) bool {
	return c.cfg().existsOnDisk()
}

func IsEmpty(c Config, empty Config) bool {
	cfg := c.cfg()
	emptyCfg := empty.cfg()
	*emptyCfg = *cfg
	return c == empty
}

func (c configData) IsDerived() bool {
	return c.origin == originDerivedLocal || c.origin == originDerivedNew
}

func IsDerived(c Config) bool {
	return c.cfg().IsDerived()
}

func (c *configData) ToDerived() *configData {
	origin := ""
	if c.origin == originLocal || c.origin == originDerivedLocal {
		origin = originDerivedLocal
	} else if c.origin == originNew || c.origin == originDerivedNew {
		origin = originDerivedNew
	}
	return &configData{
		root:   c.root,
		origin: origin,
	}
}

func Load(c Config, root string) {
	c.cfg().root = root
	c.cfg().origin = ""
	data, err := ioutil.ReadFile(Path(c))
	if err != nil {
		return
	}
	err = json.Unmarshal(data, c)
	if err != nil {
		panic(err)
	}
	c.cfg().origin = originLocal
}

func LoadDerived(c Config, root string) {
	Load(c, root)
	if c.cfg().origin == originLocal {
		c.cfg().origin = originDerivedLocal
	}
}

func Save(c Config) {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(Path(c), data, 0777)
	if err != nil {
		panic(err)
	}
}
