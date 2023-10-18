import Foundation

func hash(_ str: String, _ x: Int, _ p: Int) -> Int {
    var curX = 1
    var s = 0
    for code in str.utf8 {
        s += (Int(code) * curX) % p
        curX *= x
    }
    return s % p
}

func getOccurrences(_ pattern: String, _ text: String) -> [Int] {
    let p = 1000000007
    let x = 263
    var result = [Int]()
    let x_pm1 = Int(pow(Double(x),Double(pattern.count - 1)))
    let hp = hash(pattern, x, p)
    var cur_h: Int = hash(String(text.suffix(pattern.count)), x, p)
    if cur_h == hp {
        if String(text.suffix(pattern.count)) == pattern {
            result.append(text.count - pattern.count)
        }
    }
    var i = text.count - pattern.count - 1
    while(i >= 0) {
        if let o1 = String(Array(text)[i + pattern.count]).utf8.first {
            if let o2 = String(Array(text)[i]).utf8.first {
                cur_h = (((cur_h - (((Int(o1) * x_pm1) % p) + p) % p) + p) % p * x + Int(o2) % p) % p
                if cur_h == hp {
                    result.append(i)
                }
            }
        }
        i -= 1
    }
    return result.reversed()
}


func main() {
    print("Enter pattern")
    if let pattern = readLine() {
        print("Enter filepath")
        if let filepath = readLine() {
            if let text = try? String(contentsOfFile: filepath) {
                let res = getOccurrences(pattern, text)
                print(res)
            }
            else {
                print("File reading error")
            }
        }
    }
}

main()