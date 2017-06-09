// Code generated by go-bindata.
// sources:
// template/40x.tmpl
// template/50x.tmpl
// template/include/base.html
// template/layout/www/about.html
// template/layout/www/demo.html
// template/layout/www/home.html
// template/layout/www/index.html
// template/layout/www/login.html
// template/layout/www/register.html
// template/layout/www/user.html
// template/layout.tmpl
// template/pongo2/base.html
// template/pongo2/web/about.html
// template/pongo2/web/demo.html
// template/pongo2/web/home.html
// template/pongo2/web/index.html
// template/pongo2/web/jwt_tester.html
// template/pongo2/web/login.html
// template/pongo2/web/register.html
// template/pongo2/web/user.html
// template/pongo2/web/ws.html
// DO NOT EDIT!

package template

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// bindataRead reads the given file from disk. It returns an error on failure.
func bindataRead(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// template40xTmpl reads file data from disk. It returns an error on failure.
func template40xTmpl() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/template/40x.tmpl"
	name := "template/40x.tmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// template50xTmpl reads file data from disk. It returns an error on failure.
func template50xTmpl() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/template/50x.tmpl"
	name := "template/50x.tmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templateIncludeBaseHtml reads file data from disk. It returns an error on failure.
func templateIncludeBaseHtml() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/template/include/base.html"
	name := "template/include/base.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templateLayoutWwwAboutHtml reads file data from disk. It returns an error on failure.
func templateLayoutWwwAboutHtml() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/template/layout/www/about.html"
	name := "template/layout/www/about.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templateLayoutWwwDemoHtml reads file data from disk. It returns an error on failure.
func templateLayoutWwwDemoHtml() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/template/layout/www/demo.html"
	name := "template/layout/www/demo.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templateLayoutWwwHomeHtml reads file data from disk. It returns an error on failure.
func templateLayoutWwwHomeHtml() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/template/layout/www/home.html"
	name := "template/layout/www/home.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templateLayoutWwwIndexHtml reads file data from disk. It returns an error on failure.
func templateLayoutWwwIndexHtml() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/template/layout/www/index.html"
	name := "template/layout/www/index.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templateLayoutWwwLoginHtml reads file data from disk. It returns an error on failure.
func templateLayoutWwwLoginHtml() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/template/layout/www/login.html"
	name := "template/layout/www/login.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templateLayoutWwwRegisterHtml reads file data from disk. It returns an error on failure.
func templateLayoutWwwRegisterHtml() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/template/layout/www/register.html"
	name := "template/layout/www/register.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templateLayoutWwwUserHtml reads file data from disk. It returns an error on failure.
func templateLayoutWwwUserHtml() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/template/layout/www/user.html"
	name := "template/layout/www/user.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templateLayoutTmpl reads file data from disk. It returns an error on failure.
func templateLayoutTmpl() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/template/layout.tmpl"
	name := "template/layout.tmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatePongo2BaseHtml reads file data from disk. It returns an error on failure.
func templatePongo2BaseHtml() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/template/pongo2/base.html"
	name := "template/pongo2/base.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatePongo2WebAboutHtml reads file data from disk. It returns an error on failure.
func templatePongo2WebAboutHtml() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/template/pongo2/web/about.html"
	name := "template/pongo2/web/about.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatePongo2WebDemoHtml reads file data from disk. It returns an error on failure.
func templatePongo2WebDemoHtml() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/template/pongo2/web/demo.html"
	name := "template/pongo2/web/demo.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatePongo2WebHomeHtml reads file data from disk. It returns an error on failure.
func templatePongo2WebHomeHtml() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/template/pongo2/web/home.html"
	name := "template/pongo2/web/home.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatePongo2WebIndexHtml reads file data from disk. It returns an error on failure.
func templatePongo2WebIndexHtml() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/template/pongo2/web/index.html"
	name := "template/pongo2/web/index.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatePongo2WebJwt_testerHtml reads file data from disk. It returns an error on failure.
func templatePongo2WebJwt_testerHtml() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/template/pongo2/web/jwt_tester.html"
	name := "template/pongo2/web/jwt_tester.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatePongo2WebLoginHtml reads file data from disk. It returns an error on failure.
func templatePongo2WebLoginHtml() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/template/pongo2/web/login.html"
	name := "template/pongo2/web/login.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatePongo2WebRegisterHtml reads file data from disk. It returns an error on failure.
func templatePongo2WebRegisterHtml() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/template/pongo2/web/register.html"
	name := "template/pongo2/web/register.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatePongo2WebUserHtml reads file data from disk. It returns an error on failure.
func templatePongo2WebUserHtml() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/template/pongo2/web/user.html"
	name := "template/pongo2/web/user.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatePongo2WebWsHtml reads file data from disk. It returns an error on failure.
func templatePongo2WebWsHtml() (*asset, error) {
	path := "/Users/Steven/Develop/Go/src/echo-web/template/pongo2/web/ws.html"
	name := "template/pongo2/web/ws.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"template/40x.tmpl": template40xTmpl,
	"template/50x.tmpl": template50xTmpl,
	"template/include/base.html": templateIncludeBaseHtml,
	"template/layout/www/about.html": templateLayoutWwwAboutHtml,
	"template/layout/www/demo.html": templateLayoutWwwDemoHtml,
	"template/layout/www/home.html": templateLayoutWwwHomeHtml,
	"template/layout/www/index.html": templateLayoutWwwIndexHtml,
	"template/layout/www/login.html": templateLayoutWwwLoginHtml,
	"template/layout/www/register.html": templateLayoutWwwRegisterHtml,
	"template/layout/www/user.html": templateLayoutWwwUserHtml,
	"template/layout.tmpl": templateLayoutTmpl,
	"template/pongo2/base.html": templatePongo2BaseHtml,
	"template/pongo2/web/about.html": templatePongo2WebAboutHtml,
	"template/pongo2/web/demo.html": templatePongo2WebDemoHtml,
	"template/pongo2/web/home.html": templatePongo2WebHomeHtml,
	"template/pongo2/web/index.html": templatePongo2WebIndexHtml,
	"template/pongo2/web/jwt_tester.html": templatePongo2WebJwt_testerHtml,
	"template/pongo2/web/login.html": templatePongo2WebLoginHtml,
	"template/pongo2/web/register.html": templatePongo2WebRegisterHtml,
	"template/pongo2/web/user.html": templatePongo2WebUserHtml,
	"template/pongo2/web/ws.html": templatePongo2WebWsHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"template": &bintree{nil, map[string]*bintree{
		"40x.tmpl": &bintree{template40xTmpl, map[string]*bintree{}},
		"50x.tmpl": &bintree{template50xTmpl, map[string]*bintree{}},
		"include": &bintree{nil, map[string]*bintree{
			"base.html": &bintree{templateIncludeBaseHtml, map[string]*bintree{}},
		}},
		"layout": &bintree{nil, map[string]*bintree{
			"www": &bintree{nil, map[string]*bintree{
				"about.html": &bintree{templateLayoutWwwAboutHtml, map[string]*bintree{}},
				"demo.html": &bintree{templateLayoutWwwDemoHtml, map[string]*bintree{}},
				"home.html": &bintree{templateLayoutWwwHomeHtml, map[string]*bintree{}},
				"index.html": &bintree{templateLayoutWwwIndexHtml, map[string]*bintree{}},
				"login.html": &bintree{templateLayoutWwwLoginHtml, map[string]*bintree{}},
				"register.html": &bintree{templateLayoutWwwRegisterHtml, map[string]*bintree{}},
				"user.html": &bintree{templateLayoutWwwUserHtml, map[string]*bintree{}},
			}},
		}},
		"layout.tmpl": &bintree{templateLayoutTmpl, map[string]*bintree{}},
		"pongo2": &bintree{nil, map[string]*bintree{
			"base.html": &bintree{templatePongo2BaseHtml, map[string]*bintree{}},
			"web": &bintree{nil, map[string]*bintree{
				"about.html": &bintree{templatePongo2WebAboutHtml, map[string]*bintree{}},
				"demo.html": &bintree{templatePongo2WebDemoHtml, map[string]*bintree{}},
				"home.html": &bintree{templatePongo2WebHomeHtml, map[string]*bintree{}},
				"index.html": &bintree{templatePongo2WebIndexHtml, map[string]*bintree{}},
				"jwt_tester.html": &bintree{templatePongo2WebJwt_testerHtml, map[string]*bintree{}},
				"login.html": &bintree{templatePongo2WebLoginHtml, map[string]*bintree{}},
				"register.html": &bintree{templatePongo2WebRegisterHtml, map[string]*bintree{}},
				"user.html": &bintree{templatePongo2WebUserHtml, map[string]*bintree{}},
				"ws.html": &bintree{templatePongo2WebWsHtml, map[string]*bintree{}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

