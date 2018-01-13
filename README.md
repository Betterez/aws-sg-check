aws-sg-check
======================
What does it do:
---------------------
Checks and remove unused security groups in aws.

What won't it remove?
--------------------------
Any security group that is being used or being referenced.

License:
------------
View the license file.

Requirements
-------------------
1. **go** version 1.92 was used here. Earlier versions will probably work, but why bother? It's free. Note that go must be installed properly. Attached hereby the `install_go.sh` script (must be run as `root`) that will install version 1.92 on a Linux machine.
2. **make** will make working easier. Not a must, but there is a make file, it does simplify things
3. Run `make setup` will install dependencies.
4. If you wish to edit or remove parts, I recommend using the *atom* editor. this is also covered in the make file: run `make edit` to edit the file with atom.
If you have *go-plus* installed, you're in for a treat.
5. Running is done with `make run`.
6. Since this is pretty straight forward, there are no tests.
