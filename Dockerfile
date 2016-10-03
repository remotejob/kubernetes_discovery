FROM scratch
EXPOSE 8000

COPY cvserver /
CMD ["/cvserver"]