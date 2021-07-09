package main

import "strings"

func checkUrl(url string , c *NewRT) int{
	prefixUrlCount := len(c.prefix)
	for k, prefixUrl := range c.prefix {
		urlLen := len(url)
		prefixUrlLen := len(prefixUrl)
		if prefixUrl[prefixUrlLen-1] == '/' {
			urlLen = strings.LastIndexByte(url, '/')
			prefixUrlLen--
		}else{
			qI:= strings.LastIndexByte(url, '?')
			if  qI != -1 {
				urlLen = qI
			}
		}
		if urlLen >= prefixUrlLen {
			j := 0
			for i := 0; i < urlLen; i++ {
				if j >= prefixUrlLen {
					break
				}
				if url[i] != prefixUrl[j] {
					if prefixUrl[j] == '*' {
						for i < urlLen && url[i] != '/' {
							i++
						}
						if i == urlLen && j == prefixUrlLen-1 {
							return 1
						}
						j++
					} else if k < prefixUrlCount {
						break
					} else {
						return 0
					}
				} else if i == urlLen-1 && j == prefixUrlLen-1 {
					return 1
				}
				j++
			}
		}
	}
	return 0

}