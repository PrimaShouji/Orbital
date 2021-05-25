import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { tap } from "rxjs/operators";

@Injectable({
	providedIn: "root",
})
export class OrbitalServerService {
	private static readonly API_LOCATION = "http://localhost:7558";

	constructor(private http: HttpClient) {}

	test() {
		return this.http
			.get(this.path("/admin/test"), { responseType: "text" })
			.pipe(tap(console.log, console.error));
	}

	private path(relativePath: string): string {
		return OrbitalServerService.API_LOCATION + relativePath;
	}
}
