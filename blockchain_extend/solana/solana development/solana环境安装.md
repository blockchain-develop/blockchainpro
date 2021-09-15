# solana环境安装(mac)

## 安装solana

```
sh -c "$(curl -sSfL https://release.solana.com/v1.7.11/install)"
```

可以替换v1.7.11为其他版本，或者使用stable、beta、edge。

安装完成后有以下输出：
```
downloading v1.7.11 installer
Configuration: /home/solana/.config/solana/install/config.yml
Active release directory: /home/solana/.local/share/solana/install/active_release
* Release version: v1.7.11
* Release URL: https://github.com/solana-labs/solana/releases/download/v1.7.11/solana-release-x86_64-unknown-linux-gnu.tar.bz2
Update successful
```

安装的最后有以下输出：
```
Please update your PATH environment variable to include the solana programs:
    PATH="/Users/tangaoyuan/.local/share/solana/install/active_release/bin:$PATH"
```

安装完成后，按照上面的提示将环境变量添加到PATH中。

测试安装结果：
```
solana --version
```

如果想升级solana：
```
solana-install update
```

## 安装solana定制的rust

1. 进入项目目录
2. cargo-build-bpf
3. 后续编译solana合约可以cargo build-bpf

