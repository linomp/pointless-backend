package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type ServerMetrics struct {
	Host        string
	Timestamp   string
	CPUUsage    string
	MemoryUsage string
}

func getStatus(r *http.Request) ServerMetrics {
	// Get the client IP address
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	// Get the current timestamp in ISO 8601 format
	time := time.Now().UTC().Format(time.RFC3339)

	// Get CPU and Memory usage
	cpuUsage, _ := cpu.Percent(0, false)
	memoryUsage, _ := mem.VirtualMemory()

	return ServerMetrics{
		Host:        ip,
		Timestamp:   time,
		CPUUsage:    fmt.Sprintf("%.2f %%", cpuUsage[0]),
		MemoryUsage: fmt.Sprintf("%.2f %%", memoryUsage.UsedPercent),
	}
}

func formatTimestamp(timestamp string) string {
	return timestamp[11:19] + " (UTC)"
}

func generateServerStatusHTML(metrics ServerMetrics) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html lang=en>
<head>
	<meta charset=UTF-8>
	<meta name=viewport content=width=device-width, initial-scale=1.0>
	<style>
		@import url('https://fonts.googleapis.com/css?family=Open+Sans&display=swap');
		html {
			position: relative;
			overflow-x: hidden !important;
		}
		* {
			box-sizing: border-box;
		}
		body {
			font-family: 'Open Sans', sans-serif;
			color: #324e63;
		}
		a {
			color: inherit;
			text-decoration: inherit;
		}
		.wrapper {
			width: 100%%;
			width: 100%%;
			height: auto;
			min-height: 90vh;
			padding: 50px 20px;
			padding-top: 100px;
			display: flex;
		}
		.instance-card {
			width: 100%%;
			min-height: 380px;
			margin: auto;
			box-shadow: 12px 12px 2px 1px rgba(13, 28, 39, 0.4);
			background: #fff;
			border-radius: 15px;
			border-width: 1px;
			max-width: 500px;
			position: relative;
			border: thin groove #9c83ff;
		}
		.instance-card__cnt {
			margin-top: 35px;
			text-align: center;
			padding: 0 20px;
			padding-bottom: 40px;
			transition: all .3s;
		}
		.instance-card__name {
			font-weight: 700;
			font-size: 24px;
			color: #6944ff;
			margin-bottom: 15px;
		}
		.instance-card-inf__item {
			padding: 10px 35px;
			min-width: 150px;
		}
		.instance-card-inf__title {
			font-weight: 700;
			font-size: 27px;
			color: #324e63;
		}
		.instance-card-inf__txt {
			font-weight: 500;
			margin-top: 7px;
		}
		.secondary {
			color: #9c83ff;
		}
		.secondary:hover {
			text-decoration: underline;
		}
	</style>
	<title>😈️ pointless-status 😈</title>
</head>
<body>
	<div class=wrapper>
		<div class=instance-card>
			<div class=instance-card__cnt>
				<div class=instance-card__name>😈️ Server is running! 😈️</div>
				<div class=instance-card-inf>
					<div class=instance-card-inf__item>
						<div class=instance-card-inf__txt>Client</div>
						<div class=instance-card-inf__title>%s</div>
					</div>
					<div class=instance-card-inf__item>
						<div class=instance-card-inf__txt>Time</div>
						<div class=instance-card-inf__title>%s</div>
					</div>
					<div class=instance-card-inf__item>
						<div class=instance-card-inf__txt>CPU Usage</div>
						<div class=instance-card-inf__title>%s</div>
					</div>
					<div class=instance-card-inf__item>
						<div class=instance-card-inf__txt>Memory usage</div>
						<div class=instance-card-inf__title>%s</div>
					</div>
					<div class=instance-card-inf__item>
                        <div class="instance-card-inf__title secondary">
						  <a href="https://github.com/linomp/pointless-backend" target="_blank">
						    <svg xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" width="40" height="40" viewBox="0,0,256,256">
						      <g fill="#9c83ff" fill-rule="nonzero" stroke="none" stroke-width="1" stroke-linecap="butt" stroke-linejoin="miter" stroke-miterlimit="10" stroke-dasharray="" stroke-dashoffset="0" font-family="none" font-weight="none" font-size="none" text-anchor="none" style="mix-blend-mode: normal"><g transform="scale(8.53333,8.53333)"><path d="M15,3c-6.627,0 -12,5.373 -12,12c0,5.623 3.872,10.328 9.092,11.63c-0.056,-0.162 -0.092,-0.35 -0.092,-0.583v-2.051c-0.487,0 -1.303,0 -1.508,0c-0.821,0 -1.551,-0.353 -1.905,-1.009c-0.393,-0.729 -0.461,-1.844 -1.435,-2.526c-0.289,-0.227 -0.069,-0.486 0.264,-0.451c0.615,0.174 1.125,0.596 1.605,1.222c0.478,0.627 0.703,0.769 1.596,0.769c0.433,0 1.081,-0.025 1.691,-0.121c0.328,-0.833 0.895,-1.6 1.588,-1.962c-3.996,-0.411 -5.903,-2.399 -5.903,-5.098c0,-1.162 0.495,-2.286 1.336,-3.233c-0.276,-0.94 -0.623,-2.857 0.106,-3.587c1.798,0 2.885,1.166 3.146,1.481c0.896,-0.307 1.88,-0.481 2.914,-0.481c1.036,0 2.024,0.174 2.922,0.483c0.258,-0.313 1.346,-1.483 3.148,-1.483c0.732,0.731 0.381,2.656 0.102,3.594c0.836,0.945 1.328,2.066 1.328,3.226c0,2.697 -1.904,4.684 -5.894,5.097c1.098,0.573 1.899,2.183 1.899,3.396v2.734c0,0.104 -0.023,0.179 -0.035,0.268c4.676,-1.639 8.035,-6.079 8.035,-11.315c0,-6.627 -5.373,-12 -12,-12z"></path></g></g>
						    </svg>
					      </a>
						</div>
                        <div class="instance-card-inf__txt secondary"><a href="/oauthdemo" target="_blank">/oauthdemo</a></div>
                    </div>
				</div>
			</div>
		</div>
</body>
</html>
	`, metrics.Host, formatTimestamp(metrics.Timestamp), metrics.CPUUsage, metrics.MemoryUsage)
}
