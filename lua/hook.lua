-- require("./sys")

DIR_GITHOOK = "$GOPATH/src/github.com/sunreaver/github_hook/"

function killprog( prog )
	if type(prog) ~= "string" then
	    return
	end
	os.execute("pkill -f " .. prog)
end

function gitpull(...)
	local exc = {}
	exc[1] = "cd " .. DIR_GITHOOK
	exc[2] = "git pull"
	local result = ""
	for i=1,#exc do
		result = result .. exc[i] .. ";"
	end
	r = string.sub(result, 0, string.len(result)-1)
	-- print(r)
	os.execute(r)
end

function restart(...)
	local exc = {}
	exc[1] = "cd " .. DIR_GITHOOK
	exc[2] = "go build"
	exc[3] = "mv github_hook github_hook_run"
	exc[4] = "nohup ./github_hook_run >>nohup.out 2>&1 &"
	local result = ""
	for i=1,table.maxn(exc) do
		result = result .. exc[i] .. ";"
	end
	os.execute(string.sub(result, 0, string.len(result)-1))
end

print("pull...")
gitpull()
print("kill...")
killprog("github_hook_run")
print("restart...")
restart()
print(string.format([[↑ DO AT: %s]], os.date("%Y-%m-%d %H:%M:%S")))

