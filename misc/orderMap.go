package misc

// func MarshalJSONWithOrder(obj interface{}) ([]byte, error) {
// 	val := reflect.ValueOf(obj)
// 	typ := reflect.TypeOf(obj)

// 	var fields []fieldWithOrder

// 	for i := 0; i < val.NumField(); i++ {
// 		fieldName := typ.Field(i).Name

// 		order := getTagValue(typ.Field(i), "json", "order")
// 		fields = append(fields, fieldWithOrder{
// 			Name:  fieldName,
// 			Order: order,
// 		})
// 	}

// 	sort.Slice(fields, func(i, j int) bool {
// 		return fields[i].Order < fields[j].Order
// 	})

// 	//result := orderedmap.New[string, any]()
// 	result := orderedmap.New()
// 	for _, f := range fields {
// 		//result[f.Name] = val.FieldByName(f.Name).Interface()
// 		result.Set(f.Name, val.FieldByName(f.Name).Interface())
// 	}

// 	return json.Marshal(result)
// }

// // getTagValue
// func getTagValue(field reflect.StructField, tag string, tagField string) int {
// 	tagValue, _ := field.Tag.Lookup(tag)

// 	// Split the tag string by ","
// 	tagParts := strings.Split(tagValue, ",")

// 	// Iterate through the tag parts to find the "order" value
// 	res := -1
// 	for _, part := range tagParts {
// 		// Check if the part starts with "order:"
// 		prefix := fmt.Sprintf("%s:", tagField)
// 		if strings.HasPrefix(part, prefix) {
// 			// Extract the numeric value after "order:"
// 			orderStr := strings.TrimPrefix(part, prefix)
// 			order, err := strconv.Atoi(orderStr)
// 			if err == nil {
// 				res = order
// 			}
// 		}
// 	}

// 	return res
// }

// // fieldWithOrder
// type fieldWithOrder struct {
// 	Name  string
// 	Order int
// }
