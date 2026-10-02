# TDay - To-Do

## Features

- Add severity field
- Add color to output

## Commands

- [ ] `tday upd`
- [ ] `tday done`
- [ ] `tday rm`

Expiration

- [ ] `tday ls exp`
- [ ] `tday rm exp {#}`
- [ ] `tday done exp {#}`

Expired tasks will explicitly list the date it was created
and tasks will be sorted by the TTL
on and the time it was due at will show the overdue label
and the TTL. After the TTL the task will self-destruct.

Completion

- [ ] `tday done {#}`
- [ ] `tday ls done`
- [ ] `tday done {#, #, #, #}`

Updates

- [ ] `tday upd {#}`

Tmrw flag (maybe adding...)

- [ ] `tday cmd --tmrw`
