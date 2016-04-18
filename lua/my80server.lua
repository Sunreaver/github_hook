-- require("./sys")

DIR_MYSERVER = "$GOPATH/src/github.com/sunreaver/myserver/"

function gitpull(...)
	local exc = {}
	exc[0] = "cd " .. DIR_MYSERVER
	exc[1] = "git pull"
	local result = ""
	for i=0,table.maxn(exc) do
		result = result .. exc[i] .. ";"
	end
	r = string.sub(result, 0, string.len(result)-1)
	-- print(r)
	os.execute(r)
end

function restart(...)
	local exc = {}
	exc[0] = "cd " .. DIR_MYSERVER
	exc[1] = "go build"
	exc[2] = "cp myserver.d /etc/init.d/myserver.d"
	exc[3] = "mv myserver ~/Doc/bin/"
	exc[4] = "service myserver.d restart"
	local result = ""
	for i=0,table.maxn(exc) do
		result = result .. exc[i] .. ";"
	end
	os.execute(string.sub(result, 0, string.len(result)-1))
end

print("pull...")
gitpull()
print("restart...")
restart()

