# Module Schema

This file documents the various options (required and optional) that modules can
accept: ie. the module's schema.

*Note: other supported operating systems and architectures will be added in the future.*

*Note: not all of these options are implemented.*

## Top Level Configuration

The following are all the top level module configuration fields:

| Field | Type | Purpose | Required |
| ----- | ---- | ------- | -------- |
| `name` | string | specifiy the module name | Y |
| `version` | string | version specifies the module's version (mostly for dependency version control) | N |
| `chai-version` | string | specify the Chai version the module was created on; in the form: `Major.Minor.Build` | Y |
| `caching` | bool | enable [compilation caching](#compl-caching) | N, default = `false` |

### <a name="compl-caching"></a> Compilation Caching

**Compilation caching** is profile by which the Chai compiler will cache
precompiled object files from previous builds so that it doesn't have to
recompile code that hasn't changed.  This works very similar to Makefiles
wherein the code is recompiled if the corresponding package has been updated
since the object files were last created.

This option can save a lot of compilation time on large projects.

All cached files are placed in the `.chai` directory which is located in the
same directory that contains the output path (directory or file). 

For example, if you had a project where the output path was `out/project.exe`,
then, the `.chai` directory would be dropped in `out` like so:

```language-text
project/
    out/  
        .chai/  <-- same dir as output path
            ...
        project.exe  <-- output path
    ...
```

The `.chai` directory should generally be placed in your `.gitignore`.  Not only
will large object files be placed in it but also several other associated files,
namely, `cpi` files for each object file and possibly partial `pdb` files as
well.