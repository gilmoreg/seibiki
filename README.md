# seibiki
Auto lookup for Japanese texts

## Usage

```bash
git clone github.com/gilmoreg/seibiki
cd seibiki

make setup
# now fill in MONGODB_CONNECTION_STRING in ./build/.env
# should be mongodb://<username>:<password>@<host>:<port>/<db>

# API will be available at http://localhost:3001
make up-build

# development UI will be available at http://localhost:3000
make client
```