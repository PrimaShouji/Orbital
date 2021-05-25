import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { tap } from "rxjs/operators";

@Injectable()
export class OrbitalServer {
	private static readonly API_LOCATION = "http://localhost:7558";

	constructor(private http: HttpClient) {}

	test() {
		return this.http
			.get(this.path("/admin/test"), { responseType: "text" })
			.pipe(tap(console.log, console.error));
	}

	private path(relativePath: string): string {
		return OrbitalServer.API_LOCATION + relativePath;
	}
}
