
local lfs = require"lfs"
local DIR_WARMS = "$GOPATH/src/github.com/sunreaver/warms/"

-- 查找.go文件
function attrdir (path)
	local fNames = {}
	local i = 1
    for file in lfs.dir(path) do
        if file ~= "." and file ~= ".." then
            local f = path..'/'..file
            local attr = lfs.attributes (f)
            if type(attr) == "table" then
	            -- assert (type(attr) == "table")
	            -- if attr.mode == "directory" then
	            if attr.mode == "file" then
					if string.find(file, "^.+%.go$") then
						fNames[i] = file
						i = i + 1
					end
				end
            end
        end
    end
    return fNames
end

function gitpull(path)
	local exc = {}
	exc[1] = "cd " .. path
	exc[2] = "git pull"
	local result = ""
	for i=1,#exc do
		result = result .. exc[i] .. ";"
	end
	r = string.sub(result, 0, string.len(result)-1)
	os.execute(r)
end

function buildAndDo(fileNames, path)
	for j=1,#fileNames do
		local exc = {}
		exc[1] = "cd " .. path
		exc[2] = "go build " .. fileNames[j]
		print("build: " .. fileNames[j])
		if string.find(fileNames[j], "^huaban_warm.*%.go$") then
			exc[3] = "mv " .. string.sub(fileNames[j], 0, string.len(fileNames[j]) - 3) .. " /root/Doc/bin/huaban/huaban_warm"
			exc[4] = "service huaban restart"
		elseif string.find(fileNames[j], "^stock_warms%.go$") then
			exc[3] = "mv " .. string.sub(fileNames[j], 0, string.len(fileNames[j]) - 3) .. " /root/Doc/bin/stock/stockwarm"
			exc[4] = "cp stock.json /root/Doc/bin/stock/"
		end
		local result = ""
		for i=1,#exc do
			result = result .. exc[i] .. ";"
		end
		r = string.sub(result, 0, string.len(result)-1)
		os.execute(r)
	end
end

-- do
local rp = io.popen("echo ".. DIR_WARMS)
local realPath = rp:read("*l")
gitpull(realPath)
local fns = attrdir(realPath)
print(#fns)
-- assert(type(fns) == "table")
buildAndDo(fns, realPath)

