# Handling Relations

References
Let's get back to the Group. Ideally, a group should have members, and Charlatan allows you to reference one record from another one. You can do that with the @name notation, where name is a fixture name.

Let's add a fixed owner to the group:

```yaml
user:
    admin_1:
        first_name: 'William'
        last_name: 'Wallace'
        group_id: '@group_1'
    user_{1..10}:
        first_name: 'Bob {current}'
        last_name: 'The sponge'
        group_id: '@group_2'

group:
    group_1:
        name: admin
    group_2:
        name: reader
```