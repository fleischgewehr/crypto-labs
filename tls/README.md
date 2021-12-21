# TLS configuration

Add a TLS layer to your app or web-app. This task is pretty simple considering every modern web-framework has an easy to configure/use TLS support. You are free to use reverse proxy with TLS before your server instead of relying on framework-implementation.

Part 1.
1. Choose an appropriate TLS version for 2020.
2. Choose an appropriate and secure enough ciphersuite(s) for your server.
3. You may use self-signed certificates. If you like, you are welcome to use any CA-signed certs. But do not use automated tools like certbot unless you are familiar with how to configure everything manually. Upload all your configs and relevant code to a public repository on github.

You will be graded based on your configs and your report. In your report please describe with your own words:
Why did you choose the ciphersuites/protocol versions/algorithms etc.
Where does every component (keys, certificates) reside on your server.
Be prepared to show every step of config at request.
