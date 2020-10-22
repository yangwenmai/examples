package main

import "sort"

func uniq(list []string) []string {
	if list == nil {
		return nil
	}
	out := make([]string, len(list))
	copy(out, list)
	sort.Strings(out)
	uniq := out[:0]
	for _, x := range out {
		if len(uniq) == 0 || uniq[len(uniq)-1] != x {
			uniq = append(uniq, x)
		}
	}
	return uniq
}

func uniq2(list []string) []string {
	m := map[string]struct{}{}
	for _, l := range list {
		if _, ok := m[l]; ok {
			continue
		}
		m[l] = struct{}{}
	}
	var ret []string
	for v := range m {
		ret = append(ret, v)
	}
	return ret
}

func uniq3(list []string) []string {
	m := map[string]struct{}{}
	var ret []string
	for _, l := range list {
		if _, ok := m[l]; !ok {
			ret = append(ret, l)
			m[l] = struct{}{}
		}
	}
	return ret
}

func uniq4(list []string) []string {
	m := map[string]struct{}{}
	for _, l := range list {
		if _, ok := m[l]; ok {
			continue
		}
		m[l] = struct{}{}
	}
	out := make([]string, len(m))
	i := 0
	for v := range m {
		out[i] = v
		i++
	}
	return out
}
