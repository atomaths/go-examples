#!/usr/bin/env bash
#
# 외부 패키지 소스를 현재 프로젝트의 vendor/ 디렉토리 안에
# 두려고 할 때, 이 스크립트로 외부 소스의 커밋 리비전 관리.
# docker 방식.
#
set -e

# Downloads dependencies into vendor/ directory
if [[ ! -d vendor ]]; then
  mkdir vendor
fi
vendor_dir=${PWD}/vendor

rm_pkg_dir() {
  PKG=$1
  REV=$2
  (
    cd $vendor_dir
    if [[ -d src/$PKG ]]; then
      echo "src/$PKG already exists. Removing."
      rm -fr src/$PKG
    fi
  )
}

git_clone() {
  PKG=$1
  REV=$2
  (
    rm_pkg_dir $PKG $REV
    cd $vendor_dir && git clone http://$PKG src/$PKG
    cd src/$PKG && git checkout -f $REV && rm -fr .git
  )
}

hg_clone() {
  PKG=$1
  REV=$2
  (
    rm_pkg_dir $PKG $REV
    cd $vendor_dir && hg clone http://$PKG src/$PKG
    cd src/$PKG && hg checkout -r $REV && rm -fr .hg
  )
}

git_clone github.com/nsf/termbox-go 0d3ee6a7b6
#hg_clone code.google.com/p/go.net 84a4013f96e0
