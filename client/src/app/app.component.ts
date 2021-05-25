import { Component } from "@angular/core";
import { OrbitalServerService } from "./orbital/orbital-server.service";

@Component({
	selector: "app-root",
	templateUrl: "./app.component.html",
	styleUrls: ["./app.component.scss"],
})
export class AppComponent {
	title = "client";

	constructor(private orbital: OrbitalServerService) {
		this.orbital.test();
	}
}
