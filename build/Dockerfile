FROM aristorinjuang/go-builder AS BUILDER
COPY . .
RUN go build -o app cmd/app/main.go

FROM aristorinjuang/go-production
ENV TZ=Asia/Jakarta
COPY --from=BUILDER /app/app .
COPY --from=BUILDER /app/configs/.env.example .env
CMD ["./app"]