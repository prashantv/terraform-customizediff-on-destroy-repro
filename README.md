# Terraform destroy unexpected triggers CustomizeDiff

This is a repro for a bug in Terraform (or the SDK) that leads to
CustomizeDiff being called for a resource on destroys, even though the
documentation explicitly calls out "Destroy: No runs".

To reproduce, check out this repo, cd into it, and run the following commands:
```
# ensure the repo is in a clean state
$ git clean -fdx

# install the provider to the local plugins directory
$ make install

# apply the initial state
$ make apply

# run destroy with CUSTOMIZE_DIFF_PANIC=1 which will trigger
# a panic if CustomizeDiff is called (as it isn't expected).
$ make destroy

panic: customize_diff called with CUSTOMIZE_DIFF_PANIC=1
[...]
```

