==============================
GopherJS_ Virtual Pāli_ Keypad
==============================

Virtual Keypad/Keyboard for Pāli_ input. Written in GopherJS_.

Development Environment:

  - `Ubuntu 16.10`_
  - `Go 1.7.5`_
  - Chromium_ Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)


Install
+++++++

.. code-block:: bash

  $ go get -u github.com/siongui/gopherjs-pali-virtual-keypad


Example
+++++++

See Demo_ first. The following is key point in the code.

HTML_:

.. code-block:: html

  <input type="text" id="userinput"><br>
  <div id="keypad"></div>

Go_:

.. code-block:: go

  BindKeypad("userinput", "keypad")

See `example <example>`_ directory for complete example.


UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `golang initialize two dimensional array - Google search <https://www.google.com/search?q=golang+initialize+two+dimensional+array>`_

       `golang initialize two dimensional array - DuckDuckGo search <https://duckduckgo.com/?q=golang+initialize+two+dimensional+array>`_

       `golang initialize two dimensional array - Bing search <https://www.bing.com/search?q=golang+initialize+two+dimensional+array>`_

       `golang initialize two dimensional array - Yahoo search <https://search.yahoo.com/search?p=golang+initialize+two+dimensional+array>`_

       `golang initialize two dimensional array - Baidu search <https://www.baidu.com/s?wd=golang+initialize+two+dimensional+array>`_

       `golang initialize two dimensional array - Yandex search <https://www.yandex.com/search/?text=golang+initialize+two+dimensional+array>`_

.. _Pāli: https://en.wikipedia.org/wiki/Pali
.. _Ubuntu 16.10: http://releases.ubuntu.com/16.10/
.. _Go 1.7.5: https://golang.org/dl/
.. _Chromium: https://www.chromium.org/
.. _HTML: https://www.google.com/search?q=HTML
.. _Go: https://golang.org/
.. _GopherJS: https://github.com/gopherjs/gopherjs
.. _Demo: https://siongui.github.io/gopherjs-pali-virtual-keypad/
.. _UNLICENSE: http://unlicense.org/

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
