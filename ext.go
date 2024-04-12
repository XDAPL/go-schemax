package schemax

/*
NewExtensions initializes and returns a new instance of [Extensions].
*/
func NewExtensions() Extensions {
	return Extensions{make(extensions, 0)}
}

/*
String is a stringer method that returns a properly formatted
quoted descriptor list based on the number of string values
within the receiver instance.
*/
func (r Extensions) String() (s string) {
	if !r.IsZero() {
		var _s []string
		for k, v := range r.extensions {
			_s = append(_s, k)
			switch v.cast().Len() {
			case 0:
				continue
			case 1:
				if fst := v.index(0); len(fst) > 0 {
					_s = append(_s, sprintf("'%s'", fst))
				}
			default:
				_s = append(_s, v.cast().String())
			}
		}

		if len(_s) > 0 {
			s = join(_s, string(rune(32)))
		}
	}

	return
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r Extensions) IsZero() bool {
	return r.extensions == nil
}

/*
Set assigns the provided key and values instances to the
receiver instance, thereby specifying a new extension
value as described in RFC 4512.
*/
func (r Extensions) Set(key string, values ...string) {
	r.extensions.set(key, values...)
}

func (r *extensions) set(key string, values ...string) {
	_key := uc(key)
	if !r.exists(_key) {
		_values := newQStringList(`extensions`)
		for v := 0; v < len(values); v++ {
			_values.cast().Push(values[v])
		}
		if _values.cast().Len() > 0 {
			(*r)[_key] = _values
		}
	}
}

/*
Exists returns a Boolean value indicative of whether the
specified key exists within the receiver instance.  Case
is not significant in the matching process.
*/
func (r Extensions) Exists(key string) bool {
	return r.extensions.exists(key)
}

func (r extensions) exists(key string) (found bool) {
	_, found = r.get(key)
	return
}

/*
Get returns [QuotedStringList] and Boolean instances based
on a successful retrieval of values associated with key in
the receiver instance.  The Boolean value is indicative of
a positive match.  Case is not significant in the matching
process.
*/
func (r Extensions) Get(key string) (QuotedStringList, bool) {
	return r.extensions.get(key)
}

func (r extensions) get(key string) (values QuotedStringList, found bool) {
	values, found = r[uc(key)]
	return
}

/*
Keys returns all keys found within the underlying map instance.
*/
func (r Extensions) Keys() (keys []string) {
	for k, _ := range r.extensions {
		keys = append(keys, uc(k))
	}

	return
}

/*
Len returns the current integer length of the receiver
instance.
*/
func (r Extensions) Len() int {
	return r.extensions.len()
}

func (r extensions) len() int {
	return len(r)
}

func (r Extensions) tmplFunc() (e string) {
	if !r.IsZero() {
		e = " " + r.String()
	}
	return
}
