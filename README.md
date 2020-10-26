# go-abtest

- GetABTZone
- PushLabOutPut
- GetLabOutput
- GetGlobalIDType
- IsUsingDate
- SetCacheSyncDBFrequency

## example

```
...
targetZone := sdk.GetABTZone(project, userID, layerID, strconv.Itoa(int(time.Now().Unix()/3600/24)))
	switch true {
	case "A" == targetZone.Value:
	case "B" == targetZone.Value:
	default:
	}
```