package interesting

import "strconv"

func Flatten(m map[string]interface{}) map[string]interface{} {
	dst := map[string]interface{}{}


	for k, v := range m {
		flattern(m , dst, k )

		switch child := v.(type) {
		case map[string]interface{}:
			nm := Flatten(child)
			for nk, nv := range nm {
				o[k+"."+nk] = nv
			}
		case []interface{}:
			for i := 0; i < len(child); i++ {
				o[k+"."+strconv.Itoa(i)] = child[i]
			}
		default:
			o[k] = v
		}
	}
	return o
}

func flattern(src map[string]interface{} , dst map[string]interface{}, path ...string) {
	for _, v := range path {
		
	}
}
