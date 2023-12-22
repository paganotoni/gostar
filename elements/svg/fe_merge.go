// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg is generated from configuration file.
// Description:
// Scalable Vector Graphics (SVG) is an XML-based markup language for describing 
// two-dimensional based vector graphics 
// As such, it's a text-based, open Web standard for describing images that can be 
// rendered cleanly at any size and are designed specifically to work well with 
// other web standards including CSS, DOM, JavaScript, and SMIL 
// SVG is, essentially, to graphics what HTML is to text 
// SVG images and their related behaviors are defined in XML text files, which 
// means they can be searched, indexed, scripted, and compressed 
// Additionally, this means they can be created and edited with any text editor or 
// with drawing software 
// Compared to classic bitmapped image formats such as JPEG or PNG, SVG-format 
// vector images can be rendered at any size without loss of quality and can be 
// easily localized by updating the text within them, without the need of a 
// graphical editor to do so 
// With proper libraries, SVG files can even be localized on-the-fly. 
package svg

import(
    "fmt"
    "github.com/igrmk/treemap/v2"
)

// The <feMerge> SVG element allows filter effects to be applied concurrently 
// instead of sequentially. 
type FEMERGEElementBuilder struct {
    *ElementBuilder
}

// Create a new FEMERGEElementBuilder element.
// This will create a new element with the tag
// "feMerge" during rendering.
func FEMERGE(children ...ElementRenderer) *FEMERGEElementBuilder {
    return &FEMERGEElementBuilder{
        ElementBuilder: &ElementBuilder{
            Tag: []byte("feMerge"),
            IsSelfClosing: false,
            Descendants: children,
		},
    }
}

func (e *FEMERGEElementBuilder) Children(children ...ElementRenderer) *FEMERGEElementBuilder {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *FEMERGEElementBuilder) IfChildren(condition bool, children ...ElementRenderer) *FEMERGEElementBuilder {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *FEMERGEElementBuilder) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *FEMERGEElementBuilder {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *FEMERGEElementBuilder) Text(text string) *FEMERGEElementBuilder {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *FEMERGEElementBuilder) TextF(format string, args ...any) *FEMERGEElementBuilder {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *FEMERGEElementBuilder) Escaped(text string) *FEMERGEElementBuilder {
    e.Descendants = append(e.Descendants, Escaped(text))
    return e
}

func (e *FEMERGEElementBuilder) EscapedF(format string, args ...any) *FEMERGEElementBuilder {
    return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *FEMERGEElementBuilder) CustomData(key, value string) *FEMERGEElementBuilder {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = treemap.New[string,string]()
    }
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *FEMERGEElementBuilder) CustomDataRemove(key string) *FEMERGEElementBuilder {
	if e.CustomDataAttributes == nil {
		return e
	}
    e.CustomDataAttributes.Del(key)
	return e
}


// Specifies one or more classnames for an element (refers to a class in a style 
// sheet) 
func(e *FEMERGEElementBuilder) CLASS(s ...string) *FEMERGEElementBuilder{
    if e.DelimitedStrings == nil {
        e.DelimitedStrings = treemap.New[string,*DelimitedBuilder[string]]()
    }
    ds, ok := e.DelimitedStrings.Get("class")
    if !ok {
        ds = NewDelimitedBuilder[string](" ")
        e.DelimitedStrings.Set("class", ds)
    }
    ds.Add(s...)
    return e
}

// Remove the attribute class from the element.
func(e *FEMERGEElementBuilder) CLASSRemove(s ...string) *FEMERGEElementBuilder{
    if e.DelimitedStrings == nil {
        return e
    }
    ds, ok := e.DelimitedStrings.Get("class")
    if !ok {
        return e
    }
    ds.Remove(s ...)
    return e
}



// Specifies a unique id for an element 
func(e *FEMERGEElementBuilder) ID(s string) *FEMERGEElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("id", s)
    return e
}

// Remove the attribute id from the element.
func(e *FEMERGEElementBuilder) IDRemove(s string) *FEMERGEElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("id")
    return e
}


// Specifies an inline CSS style for an element 
func (e *FEMERGEElementBuilder) STYLE(k string, v string) *FEMERGEElementBuilder {
    if e.KVStrings == nil {
        e.KVStrings = treemap.New[string,*KVBuilder]()
    }
    kv, ok := e.KVStrings.Get("style")
    if !ok {
        kv = NewKVBuilder(":", ";")
        e.KVStrings.Set("style", kv)
    }
    kv.Add(k, v)
    return e
}

// Add the attributes in the map to the element.
func (e *FEMERGEElementBuilder) STYLEMap(m map[string]string) *FEMERGEElementBuilder {
    if e.KVStrings == nil {
        e.KVStrings = treemap.New[string,*KVBuilder]()
    }
    kv, ok := e.KVStrings.Get("style")
    if !ok {
        kv = NewKVBuilder(":", ";")
        e.KVStrings.Set("style", kv)
    }
    for k, v := range m {
        kv.Add(k, v)
    }
    return e
}

// Add pairs of attributes to the element.
func (e *FEMERGEElementBuilder) STYLEPairs(pairs ...string) *FEMERGEElementBuilder {
    if len(pairs) % 2 != 0 {
        panic("Must have an even number of pairs")
    }
    if e.KVStrings == nil {
        e.KVStrings = treemap.New[string,*KVBuilder]()
    }
    kv, ok := e.KVStrings.Get("style")
    if !ok {
        kv = NewKVBuilder(":", ";")
        e.KVStrings.Set("style", kv)
    }

    for i := 0; i < len(pairs); i += 2 {
        kv.Add(pairs[i], pairs[i+1])
    }

    return e
}


// Remove the attribute style from the element.
func (e *FEMERGEElementBuilder) STYLERemove(keys ...string) *FEMERGEElementBuilder {
    if e.KVStrings == nil {
        return e
    }
    kv, ok := e.KVStrings.Get("style")
    if !ok {
        return e
    }
    for _, k := range keys {
        kv.Remove(k)
    }
    return e
}



