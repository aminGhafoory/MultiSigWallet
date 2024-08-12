run: t css sqlc
	@go build -o bins/.out && bins/.out 
solc:
	@solc --bin --abi -o contracts/solcBuild contracts/multiSigWallet.sol
	
clean:
	@rm -rf contracts/abigenBuild && mkdir contracts/abigenBuild

abigen: clean
	@abigen --bin=contracts/solcBuild/MultiSigWallet.bin --abi=contracts/solcBuild/MultiSigWallet.abi --pkg=MultiSig --out=contracts/abigenBuild/MultiSigWallet.go

sqlc:
	@sqlc generate

gooseup:
	@cd models/sql/schema && goose postgres postgres://postgres:amin235711@amin-laptop.local:5432/MultiSig2 up

goosedown:
	@cd models/sql/schema && goose postgres postgres://postgres:amin235711@amin-laptop.local:5432/MultiSig2 down

css:
	@npx tailwindcss -i views/css/app.css -o public/styles.css --minify

t:
	templ generate
