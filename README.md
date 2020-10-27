# go-abtest

- GetABTZone
- PushLabOutPut
- GetLabOutput
- GetGlobalIDType
- IsUsingDate
- SetCacheSyncDBFrequency

## example
see example/main.go

```
...
targetZone := sdk.GetABTZone(projectID, userID, layerID, date)
	switch true {
	case "A" == targetZone.Value:
	      // Do Lab A
	case "B" == targetZone.Value:
	      // Do Lab B
	default:
	}
```