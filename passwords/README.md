# Password storage. Hashing

Your task is to create a simple client-server app (it may be just a web-app with front and back) with client authentication. You should use the best practices of credential storage: strong hash, salt, strong password requirements, honey pots etc. Everything that you find necessary for the maximum password storage security. 
You may want to reference owasp cheat sheets or any other reliable source of information to help you with this task.

Upload all your results to a public github repository.

Create front and back for user registration and login. You may use any framework you find comfortable to work with. After successful login you just need to at least output the result of login, but it’s not required to manage sessions, use cookies, or redirect users to some other page. The task is simply to check if login and password are indeed correct.
Many frameworks have these functions already implemented. Your job then is to modify the functionality in such a way that password storage is reliable.
Write a short report, in which include:
what algorithms/tools/security measures you chose and why
if applicable, what framework did you use, what were default security measures included and what changes you had to make in order to comply with best practices.

This task will be graded based on your code and report.

# Sensitive information storage

Your task is to modify your app previously created in lab 5 to store sensitive/private information of each user. This may be home address, phone number,personal photos or files etc.

The easiest way to store this info is to use the same db that is used to store user credentials. Though you are free to use any kind of storage that you may be familiar with: db, files, buckets etc. 

Let the user store information that he or anyone else can view through your app. E.g. think of OLX: anyone may get access to someone’s phone number. But if their db gets stolen - no phone number may be retrieved (at least that’s how it is supposed to work).
 
You may achieve this by storing AEAD-encrypted data and key in separate locations. E.g. data is stored in db, while the key rests in a config file with restricted access (possibly also encrypted, read about envelope encryption). For key storage you are free to utilise any available options, although I recommend using KMS provided by any major cloud provider. But again, this task is more about educational value than real security. Thus you may want to utilise simpler options, but include in your report that you are aware of security implications.

Write a short report which includes:
How did you implement your storage?
Why did you choose particular storage options/algorithms/libs etc?
What are the possible attack vectors on your system (i.e. how the stored information may be stolen)? 

Read up on best practices of cryptographic storage: OWASP crypto storage cheat sheet. 
