local str=io.read()
local t = {}
local i = 1
if string.len(str) == 0 then
    print ("Error: empty string")
else
    for w in str:gmatch("%S+") do
        t[i] = w
        i = i + 1
    end
    local filename = t[1] .. ".txt"
    local file = io.open(filename, "w")
    local length = string.len(t[1])
    for i = 2, #t do
        if string.len(t[i]) >= length then
            file:write(t[i] .. " ")
        end
    end
    file:close()
end