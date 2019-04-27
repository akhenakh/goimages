# caddy-hugo

This image contains a caddy with a git plugin enabled, but patched to support github auth token.

```
:80
tls off
root /public
git https://github.com/myuser/hugo-blog {
   path /data
   then  hugo --destination=/public  --cleanDestinationDir
   tokem XXXXXXXXXXXXX 
   interval 30
}
```
