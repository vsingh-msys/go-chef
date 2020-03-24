# Inspec tests for the client chef api go module
#

describe command('/go/src/chefapi_test/bin/clientkeys') do
  its('stderr') { should_not match(/error|no such file|cannot find|not used|undefined/) }
  its('stderr') { should match(%r{Error displaying key detail GET https://localhost/clients/clnt1/keys/default: 404}) }
  its('stdout') { should match(%r{^List initial client clnt1 keys \[\{KeyName:default Uri:https://localhost/clients/clnt1/keys/default Expired:false\}\]}) }
  its('stdout') { should match(%r{^List initial client clnt2 keys \[\{KeyName:default Uri:https://localhost/clients/clnt2/keys/default Expired:false\}\]}) }
  its('stdout') { should match(/^List initial client clnt3 keys \[\]/) }
  its('stdout') { should match(%r{^Add clnt1 key \{KeyName: Uri:https://localhost/clients/clnt1/keys/newkey Expired:false\}}) }
  its('stdout') { should match(/^List after add clnt1 keys \[\{(?=.*newkey)(?=.*default).*\}\]/) }
  its('stdout') { should match(%r{^Add clnt3 key \{KeyName: Uri:https://localhost/clients/clnt3/keys/default Expired:false\}}) }
  its('stdout') { should match(/^List after add clnt3 keys \[\{(?=.*default).*\}\]/) }
  its('stdout') { should match(/^Key detail clnt1 default \{KeyName:default/) }
  its('stdout') { should match(/^Key update output clnt1 default \{KeyName:default .*N0AIhUh7Fw1\+gQtR\+.*\}/) }
  its('stdout') { should match(/^Updated key detail clnt1 default \{KeyName:default .*N0AIhUh7Fw1\+gQtR\+.*\}/) }
  its('stdout') { should match(/^List delete result clnt1 keys \{KeyName:default .*N0AIhUh7Fw1\+gQtR\+.*\}/) }
end
