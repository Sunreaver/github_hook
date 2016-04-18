-- require("./sys")

DIR_MYSERVER = "$GOPATH/src/github.com/sunreaver/myserver/"

function gitpull(...)
	local exc = {}
	exc[1] = "cd " .. DIR_MYSERVER
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
	exc[1] = "cd " .. DIR_MYSERVER
	exc[2] = "go build"
	exc[3] = "cp myserver.d /etc/init.d/myserver.d"
	exc[4] = "mv myserver ~/Doc/bin/"
	exc[5] = "service myserver.d restart"
	local result = ""
	for i=1,#exc do
		result = result .. exc[i] .. ";"
	end
	os.execute(string.sub(result, 0, string.len(result)-1))
end

print("pull...")
gitpull()
print("restart...")
restart()

