package lru

import "testing"

/* test helper methods */
func checkCacheLength(t *testing.T, c *LRU, ln int) bool {
  llen := c.Len()
  mlen := len(c.hmap)
  if llen != mlen  {
    t.Errorf("List length (%d) doesn't match with the map length(%d)", llen, mlen)
    return false
  }

  if llen != ln  {
    t.Errorf("Cache length is %d, expected %d", llen, ln)
    return false
  }
  return true
}

func checkKeyOrder(t *testing.T, c *LRU, odr string, keys []string) bool {
  okeys, err := c.Keys(odr)
  if err != nil {
    t.Errorf("Error: %v",err)
    return false
  }
  if len(okeys) != len(keys) {
    t.Errorf("Number of keys returned by Key method is %d, expected %d", len(okeys), len(keys))
    return false
  }

  for i,k := range okeys {
    if k.(string) != keys[i] {
      t.Errorf("Key order is not correct at %d, got %s, expected %s", i, k.(string), keys[i])
      return false
    }
  }
  return true
}

/*test cases*/
func TestCache(t *testing.T) {
  c, err := newLRU(0)
  if err == nil {
    t.Errorf("Should throw error")
  }

  c, err = newLRU(2)
  if err != nil {
    t.Errorf("err: %v", err)
  }
  c.Add("Key1", "Value 1")
  checkCacheLength(t,c,1)
  c.Add("Key1", "Value 1a")
  checkCacheLength(t,c,1)
  c.Add("Key2", "Value 2")
  checkCacheLength(t,c,2)
  c.Add("Key3", "Value 3")
  checkCacheLength(t,c,2)
}

func TestOrderedKeys(t *testing.T) {
  c, _ := newLRU(3)

  c.Add("Key1", "Value 1")
  c.Add("Key2", "Value 2")
  c.Add("Key1", "Value 1a")
  c.Add("Key3", "Value 3")
  c.Fetch("Key2")

  _, err := c.Keys("")
  if err == nil {
    t.Error("Keys method should throw error if unsupported order directive is used")
  }

  checkKeyOrder(t,c,"asc",[]string{"Key2", "Key3", "Key1"})
  checkKeyOrder(t,c,"desc",[]string{"Key1", "Key3", "Key2"})
}

func TestKeyExist(t *testing.T) {
  c, _ := newLRU(3)

  c.Add("Key1", "Value 1")
  c.Add("Key2", "Value 2")
  c.Add("Key1", "Value 1a")
  ok := c.Exist("Key2")
  if !ok {
    t.Error("c.Exist(\"Key2\") return false, expected true")
  }
  c.Add("Key3", "Value 3")
  checkKeyOrder(t,c,"asc",[]string{"Key3", "Key1", "Key2"})

  ok = c.Exist("Key4")
  if ok {
    t.Error("c.Exist(\"Key4\") return true, expected false")
  }
}

func TestFetch(t *testing.T) {
  c, _ := newLRU(3)

  c.Add("Key1", "Value 1")
  c.Add("Key2", "Value 2")
  c.Add("Key3", "Value 3")
  c.Add("Key1", "Value 1a")
  
  val, ok := c.Fetch("Key2")
  if val == nil || !ok {
    t.Error("Fetch return false, expected true")
  }
  if val != "Value 2" {
    t.Errorf("Fetched value is %s , expected Value 2", val.(string))
  }
  checkKeyOrder(t,c,"asc",[]string{"Key2", "Key1", "Key3"})

  c.Add("Key4", "Value 4")
  val, ok = c.Fetch("Key3")
  if val !=nil || ok {
    t.Error("Fech shouldn't return value for Key3")
  }
}

func TestRemove(t *testing.T) {
  c, _ := newLRU(3)

  c.Add("Key1", "Value 1")
  c.Add("Key2", "Value 2")
  c.Add("Key3", "Value 3")
  c.Add("Key4", "Value 4")

  ok := c.Remove("Key1")
  if ok {
    t.Error("Remove should not delete a non exsiting key")
  }

  ok = c.Remove("Key3")
  if !ok {
    t.Error("Remove should not delete an exsiting key")
  }
  if c.Len() != 2 {
    t.Errorf("Expected 2 lement in catche, got %d", c.Len())
  }
}