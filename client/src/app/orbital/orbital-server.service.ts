import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";

@Injectable({
	providedIn: "root",
})
export class OrbitalServerService {
	private static readonly API_LOCATION = "http://localhost:7558";

	constructor(private http: HttpClient) {}

	isAuthenticated() {
		return this.http.get<AuthenticationState>(this.path("/authenticated"));
	}

	private path(relativePath: string): string {
		return OrbitalServerService.API_LOCATION + relativePath;
	}
}

interface AuthenticationState {
	is_authenticated: boolean;
}
