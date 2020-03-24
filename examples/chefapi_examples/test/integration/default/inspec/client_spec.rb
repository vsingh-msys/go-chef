# Inspec tests for the client chef api go module
#

describe command('/go/src/chefapi_test/bin/client') do
  its('stderr') { should match(%r{^Issue creating client: POST https://localhost/clients: 409}) }
  its('stderr') { should_not match(/error|no such file|cannot find|not used|undefined/) }
  its('stdout') { should match(%r{^List initial clients map\[(?=.*pivotal:https://localhost/clients/pivotal).*\] EndInitialList}) }
  # might want a multi line match here to test for expirationdate, key uri and privatekey
  its('stdout') { should match(%r{^Add clnt1 \{Uri:https://localhost/clients/clnt1 ChefKey:\{Name:default PublicKey:-----BEGIN}) }
  its('stdout') { should match(%r{^Add clnt2 \{Uri:https://localhost/clients/clnt2 ChefKey:\{Name:default PublicKey:-----BEGIN}) }
  its('stdout') { should match(%r{^Add clnt3 \{Uri:https://localhost/clients/clnt3 ChefKey:\{Name: PublicKey: ExpirationDate: Uri: PrivateKey:\}\}}) }
  its('stdout') { should match(%r{^Filter clients map\[clnt1:https://localhost/clients/clnt1\]}) }
  its('stdout') { should match(/^Verbose out map\[(?=.*pivotal:)/) }
  its('stdout') { should match(/^Get clnt1 \{(?=.*UserName:clnt1)(?=.*DisplayName:User1 Fullname)(?=.*Email:client1@domain.io)(?=.*ExternalAuthenticationUid:)(?=.*FirstName:client1)(?=.*LastName:fullname)(?=.*MiddleName:)(?=.*Password:)(?=.*PublicKey:)(?=.*RecoveryAuthenticationEnabled:false).*/) }
  its('stdout') { should match(/^Pivotal client \{(?=.*UserName:pivotal)(?=.*DisplayName:Chef Server Superclient)(?=.*Email:root@localhost.localdomain)(?=.*ExternalAuthenticationUid:)(?=.*FirstName:Chef)(?=.*LastName:Server)(?=.*MiddleName:)(?=.*Password:)(?=.*PublicKey:)/) }
  its('stdout') { should match(%r{^List after adding map\[(?=.*pivotal:https://localhost/clients/pivotal)(?=.*clnt1:https://localhost/clients/clnt1).*\] EndAddList}) }
  its('stdout') { should match(/^Get clnt1 \{(?=.*UserName:clnt1)(?=.*DisplayName:User1 Fullname)(?=.*Email:client1@domain.io)(?=.*ExternalAuthenticationUid:)(?=.*FirstName:client1)(?=.*LastName:fullname)(?=.*MiddleName:)(?=.*Password:)(?=.*PublicKey:)/) }
  its('stdout') { should match(%r{^List after adding map\[(?=.*pivotal:https://localhost/clients/pivotal)(?=.*clnt1:https://localhost/clients/clnt1).*\] EndAddList}) }
  # TODO: - update and create new private key
  # TODO - is admin a thing
  its('stdout') { should match(%r{^Update clnt1 partial update \{Uri:https://localhost/clients/clnt1 ChefKey:\{}) }
  its('stdout') { should match(/^Get clnt1 after partial update \{(UserName:clnt1)(?=.*DisplayName:clnt1)(?=.*Email:myclient@samp.com)(?=.*ExternalAuthenticationUid:)(?=.*FirstName:client1)(?=.*LastName:fullname)(?=.*MiddleName:)(?=.*Password:)(?=.*PublicKey:)(?=.*RecoveryAuthenticationEnabled:false).*\}/) }
  its('stdout') { should match(%r{^Update clnt1 full update \{Uri:https://localhost/clients/clnt1 ChefKey:\{Name: PublicKey: ExpirationDate: Uri: PrivateKey:\}}) }
  its('stdout') { should match(/^Get clnt1 after full update \{(UserName:clnt1)(?=.*DisplayName:clnt1)(?=.*Email:myclient@samp.com)(?=.*ExternalAuthenticationUid:)(?=.*FirstName:client)(?=.*LastName:name)(?=.*MiddleName:mid)(?=.*Password:)(?=.*PublicKey:)(?=.*RecoveryAuthenticationEnabled:false).*\}/) }
  its('stdout') { should match(%r{^Update clnt1 rename \{Uri:https://localhost/clients/clnt1new ChefKey:\{.*\}}) }
  its('stdout') { should match(/^Get clnt1 after rename \{(UserName:clnclntew)(?=.*DisplayName:clnt1)(?=.*Email:myclient@samp.com)(?=.*ExternalAuthenticationUid:)(?=.*FirstName:client)(?=.*LastName:name)(?=.*MiddleName:mid Password:)(?=.*PublicKey:)(?=.*RecoveryAuthenticationEnabled:false).*\}/) }
  its('stdout') { should match(%r{^Delete clnt1 DELETE https://localhost/clients/clnt1: 404}) }
  its('stdout') { should match(/^Delete clnt1new <nil>/) }
  its('stdout') { should match(%r{^List after cleanup map\[(?=.*pivotal:https://localhost/clients/pivotal).*\] EndCleanupList}) }
end
