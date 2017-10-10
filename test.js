;(function() {
  var pkg = {}
  pkg['github.com/matthewmueller/golly/js'] = (function() {
    function Raw(src) {}
    return {
      Raw: Raw
    }
  })()
  pkg[
    'github.com/matthewmueller/golly/testdata/22-basic-bindings'
  ] = (function() {
    var js = pkg['github.com/matthewmueller/golly/js']
    function main() {
      var $arr = definition('5', 3),
        arr = $arr[0],
        err = $arr[1]
      if (err != null) {
        console.log(err)
        return
      }
      console.log(arr)
    }
    function definition(a, b) {
      var arr = []
      var err = null
      arr = arr.concat('5')
      var c = String(parseInt(a) + b)
      arr.push(c)
      arr = arr.concat('9')
      return [arr, err]
    }
    return {
      main: main
    }
  })()
  pkg['github.com/matthewmueller/golly/testdata/22-basic-bindings'].main()
})()