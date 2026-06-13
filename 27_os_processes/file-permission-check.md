# File permission check

## Live interview task
Create a credentials file readable and writable only by its owner, and reject an existing file with broader permissions.

## Interview notes / pitfalls
- Use `OpenFile` with `O_CREATE|O_EXCL` where replacement is unsafe.
- Requested modes are affected by umask.
- Unix permission bits do not model Windows ACLs.
- Avoid time-of-check/time-of-use races when security depends on the check.
