```go
// 方案1： 过滤selected为false的商品规格
specProps := product.SpecProps
spec := models.Spec{}
err := json.Unmarshal([]byte(specProps), &spec)
if err == nil {
    sItem := make([]models.SpecChildren, 0)
    for _, v := range spec.Spec {
        scItem := make([]models.SpecChildrenItem, 0)
        for _, vv := range v.Children {
            if vv.Selected {
                vv.Selected = false
                scItem = append(scItem, vv)
            }
        }
        v.Children = scItem
        sItem = append(sItem, v)
    }
    spec.Spec = sItem

    data, err := json.Marshal(spec)
    if err == nil {
        product.SpecProps = string(data)
    }
}

// 方案2：过滤selected为false的商品规格
specProps := product.SpecProps
spec := models.Spec{}
err := json.Unmarshal([]byte(specProps), &spec)
if err == nil {
    for j := 0; j < len(spec.Spec); j++ {
        v := &spec.Spec[j] // Note: need pointer
        for i := len(v.Children) - 1; i >= 0; i-- {
            if !v.Children[i].Selected {
                // https://vbauerster.github.io/2017/04/removing-items-from-a-slice-while-iterating-in-go/
                v.Children = append(v.Children[:i], v.Children[i+1:]...)
            }
        }
    }

    specData, err := json.Marshal(spec)
    if err == nil {
        product.SpecProps = string(specData)
    }
}
```