apiVersion: apps/v1
kind: Deployment
metadata:
  name: httpclient
  labels:
    app: httpclient
spec:
  replicas: 3
  selector:
    matchLabels:
      app: httpclient
  template:
    metadata:
      labels:
        app: httpclient
    spec:
      containers:
        - name: app
          image: IMG_PLACEHOLDER
          ports:
            - containerPort: TARGET_PORT_PLACEHOLDER
              name: http-app
          resources:
            limits:
              memory: "200Mi"
              cpu: "100m"
            requests:
              memory: "100Mi"
              cpu: "50m"
          livenessProbe:
            httpGet:
              path: /ip
              port: TARGET_PORT_PLACEHOLDER
            failureThreshold: 3
            periodSeconds: 10
          readinessProbe:
            httpGet:
              path: /healthz
              port: TARGET_PORT_PLACEHOLDER
            failureThreshold: 3
            periodSeconds: 10
      # imagePullSecrets:
      #   - name: dockerhub
