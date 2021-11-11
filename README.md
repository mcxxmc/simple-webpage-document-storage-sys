# simple-webpage-document-storage-sys


(temporary intro)


A simple implementation to simulate a database system maintaining documents (instead of tables).


Backend supported by go and includes a gin server. Frontend is currently developed in React.


Unlike the file explorer in most systems, the hierarchy of the files & directories are not reflected in the physical level.


All the files owned by the same user are put under the same "physical" folder and their relations are managed by the profile of that user. All the directories are logical and virtual as they do not physically exist. Each user profile is a JSON file containing the relative relationships of all the real files and virtual directories.


User profiles are maintained by a bigger index (also a JSON file). Adding or deleting users will modify that big index, but register and unregister will not.


The two important parts of the backend is manager and cache. "manager" maintains the user collections (the profiles) for login user only. "cache" keeps all the basic user info (name, id, password, etc.) regardless of logging in or not. I am considering that replacing "cache" with an outer SQL server or Redis may simplify the system.
