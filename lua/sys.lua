-- DIR_GITHOOK = "$GOPATH/src/github.com/sunreaver/github_hook/"


function killprog( prog )
	if type(prog) ~= "string" then
	    return
	end
	os.execute("pkill -f " .. prog)
end