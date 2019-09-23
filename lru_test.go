package lru

import "testing"

/* test helper methods */
func checkCacheLength(t *testing.T, c *baseLRU, ln int) bool {
	llen := c.len()
	mlen := len(c.hmap)
	if llen != mlen {
		t.Errorf("list length (%d) doesn't match with the map length(%d)", llen, mlen)
		return false
	}

	if llen != ln {
		t.Errorf("Cache length is %d, expected %d", llen, ln)
		return false
	}
	return true
}

func checkKeyOrder(t *testing.T, c *baseLRU, odr string, keys []string) bool {
	okeys, err := c.keys(odr)
	if err != nil {
		t.Errorf("Error: %v", err)
		return false
	}
	if len(okeys) != len(keys) {
		t.Errorf("Number of keys returned by Key method is %d, expected %d", len(okeys), len(keys))
		return false
	}

	for i, k := range okeys {
		if k.(string) != keys[i] {
			t.Errorf("Key order is not correct at %d, got %s, expected %s", i, k.(string), keys[i])
			return false
		}
	}
	return true
}

/*test cases*/
func TestCache(t *testing.T) {
	c, err := newBaseLRU(0)
	if err == nil {
		t.Errorf("Should throw error")
	}

	c, err = newBaseLRU(2)
	if err != nil {
		t.Errorf("err: %v", err)
	}
	c.add("Key1", "Value 1")
	checkCacheLength(t, c, 1)
	c.add("Key1", "Value 1a")
	checkCacheLength(t, c, 1)
	c.add("Key2", "Value 2")
	checkCacheLength(t, c, 2)
	c.add("Key3", "Value 3")
	checkCacheLength(t, c, 2)
}

func TestOrderedKeys(t *testing.T) {
	c, _ := newBaseLRU(3)

	c.add("Key1", "Value 1")
	c.add("Key2", "Value 2")
	c.add("Key1", "Value 1a")
	c.add("Key3", "Value 3")
	c.fetch("Key2")

	_, err := c.keys("")
	if err == nil {
		t.Error("Keys method should throw error if unsupported order directive is used")
	}

	checkKeyOrder(t, c, "asc", []string{"Key2", "Key3", "Key1"})
	checkKeyOrder(t, c, "desc", []string{"Key1", "Key3", "Key2"})
}

func TestKeyexist(t *testing.T) {
	c, _ := newBaseLRU(3)

	c.add("Key1", "Value 1")
	c.add("Key2", "Value 2")
	c.add("Key1", "Value 1a")
	ok := c.exist("Key2")
	if !ok {
		t.Error("c.exist(\"Key2\") return false, expected true")
	}
	c.add("Key3", "Value 3")
	checkKeyOrder(t, c, "asc", []string{"Key3", "Key1", "Key2"})

	ok = c.exist("Key4")
	if ok {
		t.Error("c.exist(\"Key4\") return true, expected false")
	}
}

func TestFetch(t *testing.T) {
	c, _ := newBaseLRU(3)

	c.add("Key1", "Value 1")
	c.add("Key2", "Value 2")
	c.add("Key3", "Value 3")
	c.add("Key1", "Value 1a")

	val, ok := c.fetch("Key2")
	if val == nil || !ok {
		t.Error("Fetch return false, expected true")
	}
	if val != "Value 2" {
		t.Errorf("Fetched value is %s , expected Value 2", val.(string))
	}
	checkKeyOrder(t, c, "asc", []string{"Key2", "Key1", "Key3"})

	c.add("Key4", "Value 4")
	val, ok = c.fetch("Key3")
	if val != nil || ok {
		t.Error("Fech shouldn't return value for Key3")
	}
}

func Testremove(t *testing.T) {
	c, _ := newBaseLRU(3)

	c.add("Key1", "Value 1")
	c.add("Key2", "Value 2")
	c.add("Key3", "Value 3")
	c.add("Key4", "Value 4")

	ok := c.remove("Key1")
	if ok {
		t.Error("Remove should not delete a non exsiting key")
	}

	ok = c.remove("Key3")
	if !ok {
		t.Error("Remove should not delete an exsiting key")
	}
	if c.len() != 2 {
		t.Errorf("Expected 2 lement in catche, got %d", c.len())
	}
}
