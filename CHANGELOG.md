## [1.2.2](https://github.com/thomasklinger1234/semanticrelease/compare/v1.2.1...v1.2.2) (2023-04-28)


### Bug Fixes

* upload whole chart folder ([f105ff2](https://github.com/thomasklinger1234/semanticrelease/commit/f105ff289d82ba8632177439344b0f3f248be93d))

## [1.2.1](https://github.com/thomasklinger1234/semanticrelease/compare/v1.2.0...v1.2.1) (2023-04-27)


### Bug Fixes

* release image after publishing and executing semantic release ([1692f3f](https://github.com/thomasklinger1234/semanticrelease/commit/1692f3ff9ee4b9b8c6edce22a9755e61bf3064fe))

# [1.2.0](https://github.com/thomasklinger1234/semanticrelease/compare/v1.1.0...v1.2.0) (2023-04-27)


### Bug Fixes

* install helm repos first ([49360c4](https://github.com/thomasklinger1234/semanticrelease/commit/49360c438e29c36971b9f29b95074045a5782ae6))
* suspend helm build actions due to missing repos ([3ac7ea2](https://github.com/thomasklinger1234/semanticrelease/commit/3ac7ea24e5667782f24ef0de3d68a3283d136ab4))
* update helm repos first ([dd853a5](https://github.com/thomasklinger1234/semanticrelease/commit/dd853a5b12bb9f956daba23a714e95d4ba15421f))
* use direct path to helm charts as find does not work ([e5a035b](https://github.com/thomasklinger1234/semanticrelease/commit/e5a035bed99d0017b354552968d62ce33d27de46))


### Features

* add helm chart ([9e078fb](https://github.com/thomasklinger1234/semanticrelease/commit/9e078fb035c14d13ca6091d3549aa6826950f0d6))

# [1.1.0](https://github.com/thomasklinger1234/semanticrelease/compare/v1.0.0...v1.1.0) (2023-04-27)


### Features

* push image to ghcr ([8ed33df](https://github.com/thomasklinger1234/semanticrelease/commit/8ed33df8fac4b46c0bf12fdda496d8c2190437dc))

# 1.0.0 (2023-04-27)


### Bug Fixes

* do not run cst with TTY in GitHub actions ([49279da](https://github.com/thomasklinger1234/semanticrelease/commit/49279da26bce7f287c4664bf79be6b57d3a65d19))
* escape bad substitution ([dfe052a](https://github.com/thomasklinger1234/semanticrelease/commit/dfe052a3c93fd3d72a5243474cd9aaceeae608a1))
* grant write permissions to release workflow ([98782f2](https://github.com/thomasklinger1234/semanticrelease/commit/98782f281015ab7b6e4361cc550db64ebfb6003b))
* mount docker daemon and config for cst ([998d68d](https://github.com/thomasklinger1234/semanticrelease/commit/998d68db5a29b0d31c59a3106aa08c828ed1b2db))
* only run dependency review for prs ([2417e86](https://github.com/thomasklinger1234/semanticrelease/commit/2417e86544604f812abebb7442da11bb622b1b6f))
* use correct yaml syntax ([f0b0550](https://github.com/thomasklinger1234/semanticrelease/commit/f0b0550dece5983019124949f89e1a6d1bf37b9f))


### Features

* add codeql analysis (scheduled) ([403f001](https://github.com/thomasklinger1234/semanticrelease/commit/403f0019b05264f55c56ee9c11665c9955a4e3f2))
* add dependency review in ci pipeline ([a5d9685](https://github.com/thomasklinger1234/semanticrelease/commit/a5d9685684239dab06bda6470c77b5a27879cf72))
* add initial release workflow steps with semantic-release ([f80dc31](https://github.com/thomasklinger1234/semanticrelease/commit/f80dc311daf2daa0fda521876ffb105905d33cb4))
* upload trivy results to github ([077c1b4](https://github.com/thomasklinger1234/semanticrelease/commit/077c1b4fa1d430a40a7655325f1eff0afb22ded1))
