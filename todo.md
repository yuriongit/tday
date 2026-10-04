# TDay - To-Do

## Features

- Add severity field
- Add color to all output
- Add task type options (dev/hw/reg/sys), depending on the type the user selects,
  each type will prompt the user with
  specific fields tailored to each task type.

## Commands

- [ ] `tday upd`
- [ ] `tday done`
- [ ] `tday rm`

---

### Expiration

- [ ] `tday ls exp`
- [ ] `tday rm exp {#}`
- [ ] `tday done exp {#}`
- When expired tasks are outputted: the creation date will be included, TTL,
  alongside it's original due date (`due_at` field). Expired tasks will be
  ordered by it's TTL.

### Completion

- [ ] `tday done {#}`
- [ ] `tday ls done`

<!--- [ ] `tday done {# # # #}`-->

### Updates

- [ ] `tday upd {#}`
