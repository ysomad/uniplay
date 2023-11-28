local claims = std.extVar('claims');
{
  identity: {
    traits: {
      // VK doesn't provide an 'email_verified field'.
      //
      // Email might be empty if the user isn't allowed the 'email' scope.
      [if 'email' in claims then 'email' else null]: claims.email,
    },
  },
}