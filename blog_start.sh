#!/bin/bash

# 博客服务快速启动脚本

echo "========================================="
echo "  技术博客平台后端 - 快速启动脚本"
echo "========================================="
echo ""

# 颜色定义
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# 检查Go环境
echo "1. 检查Go环境..."
if command -v go &> /dev/null; then
    echo -e "${GREEN}✓ Go已安装: $(go version)${NC}"
else
    echo -e "${RED}✗ Go未安装,请先安装Go 1.19+${NC}"
    exit 1
fi

# 检查MySQL
echo "2. 检查MySQL服务..."
if mysql --version &> /dev/null; then
    echo -e "${GREEN}✓ MySQL客户端已安装${NC}"
    if mysql -u root -p985211 -e "SELECT 1" &> /dev/null; then
        echo -e "${GREEN}✓ MySQL服务正在运行${NC}"
    else
        echo -e "${YELLOW}⚠ MySQL服务可能未启动或密码错误${NC}"
    fi
else
    echo -e "${YELLOW}⚠ MySQL客户端未安装${NC}"
fi

# 检查Redis
echo "3. 检查Redis服务..."
if redis-cli --version &> /dev/null; then
    echo -e "${GREEN}✓ Redis客户端已安装${NC}"
    if redis-cli -h 43.143.44.118 -p 9898 -a 192047 ping &> /dev/null; then
        echo -e "${GREEN}✓ Redis服务正在运行${NC}"
    else
        echo -e "${YELLOW}⚠ Redis服务可能未启动或连接配置错误${NC}"
    fi
else
    echo -e "${YELLOW}⚠ Redis客户端未安装${NC}"
fi

echo ""
echo "========================================="
echo "  启动服务选项"
echo "========================================="
echo "1) 启动博客gRPC服务 (端口8889)"
echo "2) 启动主业务服务 (端口8887)"
echo "3) 启动API网关 (端口8888)"
echo "4) 启动所有服务"
echo "5) 初始化数据库"
echo "6) 生成Proto文件"
echo "0) 退出"
echo ""

read -p "请选择操作 (0-6): " choice

case $choice in
    1)
        echo "启动博客gRPC服务..."
        cd /Users/xiajia/Desktop/arrayDanceBackEnd/base-service/cmd/blog
        go run main.go
        ;;
    2)
        echo "启动主业务服务..."
        cd /Users/xiajia/Desktop/arrayDanceBackEnd/base-service/cmd
        go run main.go
        ;;
    3)
        echo "启动API网关..."
        cd /Users/xiajia/Desktop/arrayDanceBackEnd/douyin-api/cmd
        go run main.go
        ;;
    4)
        echo "将在新终端窗口中启动所有服务..."

        # macOS Terminal
        if [[ "$OSTYPE" == "darwin"* ]]; then
            # 终端1: 博客服务
            osascript -e 'tell app "Terminal"
                do script "cd /Users/xiajia/Desktop/arrayDanceBackEnd/base-service/cmd/blog && go run main.go"
            end tell'

            sleep 1

            # 终端2: 主服务
            osascript -e 'tell app "Terminal"
                do script "cd /Users/xiajia/Desktop/arrayDanceBackEnd/base-service/cmd && go run main.go"
            end tell'

            sleep 1

            # 终端3: API网关
            osascript -e 'tell app "Terminal"
                do script "cd /Users/xiajia/Desktop/arrayDanceBackEnd/douyin-api/cmd && go run main.go"
            end tell'

            echo -e "${GREEN}✓ 已在新终端窗口中启动所有服务${NC}"
        else
            echo "请手动在三个终端中分别启动服务"
            echo "终端1: cd base-service/cmd/blog && go run main.go"
            echo "终端2: cd base-service/cmd && go run main.go"
            echo "终端3: cd douyin-api/cmd && go run main.go"
        fi
        ;;
    5)
        echo "初始化数据库..."
        echo "请手动执行以下命令:"
        echo "mysql -u root -p douyin < /Users/xiajia/Desktop/arrayDanceBackEnd/blog-sql/init_blog.sql"
        ;;
    6)
        echo "生成Proto文件..."
        cd /Users/xiajia/Desktop/arrayDanceBackEnd/base-service/proto/blog

        # 检查protoc是否安装
        if ! command -v protoc &> /dev/null; then
            echo -e "${RED}✗ protoc未安装${NC}"
            echo "请执行: brew install protobuf (macOS)"
            echo "或: sudo apt-get install protobuf-compiler (Linux)"
            exit 1
        fi

        # 检查protoc-gen-go是否安装
        if ! command -v protoc-gen-go &> /dev/null; then
            echo "安装protoc-gen-go..."
            go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
        fi

        # 检查protoc-gen-go-grpc是否安装
        if ! command -v protoc-gen-go-grpc &> /dev/null; then
            echo "安装protoc-gen-go-grpc..."
            go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
        fi

        # 生成代码
        echo "生成Go代码..."
        protoc --proto_path=. \
          --go_out=paths=source_relative:. \
          --go-grpc_out=paths=source_relative:. \
          blog.proto

        if [ $? -eq 0 ]; then
            echo -e "${GREEN}✓ Proto文件生成成功${NC}"
            ls -lh blog.pb.go blog_grpc.pb.go
        else
            echo -e "${RED}✗ Proto文件生成失败${NC}"
        fi
        ;;
    0)
        echo "退出"
        exit 0
        ;;
    *)
        echo -e "${RED}无效的选择${NC}"
        exit 1
        ;;
esac
