# dexcom-go
This package serves as an SDK that makes it easy to interact with the Dexcom API. We make use of a simple design,
exposing the most critical data in structs that are easy to consume.

Supported Methods:
- [x] Calibrations
- [x] Data Range
- [x] Devices
- [x] Estimated Glucose Values
- [ ] Events
- [ ] Statistics

Please feel free to make a pull request to support more methods!

## Example
```go
c := dexcom.New("<your-access-token>")
egvs, err := c.GetEGVs(time.Now().Add(-1 * time.Hour), time.Now())
if err != nil {
	panic(err)
}

fmt.Println(egvs)
```

## Sandbox Mode
To test with Sandbox data, simply overwrite the `Endpoint` field on the `Client`. An example:
```go
c := dexcom.New("<your-access-token>")
c.Endpoint = "https://sandbox-api.dexcom.com"
```

## Debugging
Sometimes, it is desirable to have logs for each request, and logs detailing errors when they happen. To log information
from this package, implement the `Logger`, and overwrite the `Logger` field on the `Client`. An example:
```go
c := dexcom.New("<your-access-token>")
c.Logger = myBeautifulLogger{}
```

Obviously, you should not panic in a real-life scenario. This example will return all of your estimated glucose values
for the past hour.

## Contributing
- [ ] Better error handling (status code checks, logging, etc)

## Testing
We use an environment variable `TEST_ACCESS_TOKEN` to run our integration tests. To run the test suite, this variable 
**must be set**. To run the tests:
```sh
go test -v github.com/t1labs/dexcom-go
```

## License

Copyright © 2019 t1labs

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this work except in compliance with the License. You may obtain a copy of the License in the LICENSE file, or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.