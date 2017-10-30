# Golink Redirector

## Background

I had first encounted golinks at my time at Google, working as a Software 
Engineer, Tools and Infrastructure. It is a link redirector but with specific 
names as links. I had thought that having this as a personal service with my 
own links would be a great tool to use when trying to remember obscure links 
and be easier to share.

This application is written to be deployed via Google Cloud Platform and uses 
AppEngine and Datastore.

You can view a live version of this code @ [go.noxoin.com](https://go.noxoin.com/).

## Development

### Datastore Emulation

During development and testing, you will need to have an emulated datastore by 
following instructions [here](https://cloud.google.com/appengine/docs/standard/go/tools/using-local-server).

```
# Start emulated datastore for dev environment
gcloud beta emulators datastore start
```

Also, remember to set your environment variable on each terminal that runs the 
application.

```
# Have application pointing to dev environment and not prod datastore
$(gcloud beta emulators datastore env-init)
```


## License

	The code in this repository is licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.

**NOTE**: This software depends on other packages that may be licensed under different open source licenses.
