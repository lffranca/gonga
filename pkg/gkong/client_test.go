package gkong

//func TestServiceService_List(t *testing.T) {
//	ctx := context.Background()
//
//	gatewayURL := "http://localhost:8001"
//	client, err := New(&gatewayURL, nil)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	size := 2
//	offset := "WyIzMjk2NjM5Mi00ZmUwLTRjY2ItYjE0My0xNmZlNjI5MDQ1NDMiXQ"
//	items, options, err := client.Service.List(ctx, &size, &offset, nil, nil)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	jsonResult, _ := json.Marshal(map[string]interface{}{
//		"items":   items,
//		"options": options,
//	})
//	log.Println(string(jsonResult))
//}
