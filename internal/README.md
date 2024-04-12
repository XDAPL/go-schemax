## Internal Packages

The `internal` directory contains all of the definitions included with this package, organized within subdirectories named after the standards from which they originate.

While this "library" contains a large number of definitions, this is by no means a complete manifest.  Many vendors incorporate custom definitions into their directory products which may not originate from any standard.

All definitions have been modified in the following manners:

  - Ordering of fields per RFC 4512 -- for instance, NAME before DESC -- is enforced
  - Include `X-ORIGIN` values which identify standard(s) of origin

