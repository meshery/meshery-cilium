<p style="text-align:center;" align="center"><a href="https://layer5.io/meshery"><picture align="center">
  <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/meshery/meshery/master/.github/assets/images/meshery/meshery-logo-dark-text-side.svg"  width="70%" align="center" style="margin-bottom:20px;">
  <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/meshery/meshery/master/.github/assets/images/meshery/meshery-logo-light-text-side.svg" width="70%" align="center" style="margin-bottom:20px;">
  <img alt="Shows an illustrated light mode meshery logo in light color mode and a dark mode meshery logo dark color mode." src="https://raw.githubusercontent.com/meshery/meshery/master/.github/assets/images/meshery/meshery-logo-tag-light-text-side.png" width="70%" align="center" style="margin-bottom:20px;">
</picture></a><br /><br /></p>
<p align="center">
 
<div align="center">

[![Docker Pulls](https://img.shields.io/docker/pulls/layer5/meshery-cilium.svg)](https://hub.docker.com/r/layer5/meshery-cilium)
[![Go Report Card](https://goreportcard.com/badge/github.com/layer5io/meshery-cilium)](https://goreportcard.com/report/github.com/layer5io/meshery-cilium)
<a href="https://github.com/meshery/meshery-cilium/actions" alt="Build Status">
  <img src="https://img.shields.io/github/workflow/status/meshery/meshery-cilium/Meshery%20Adapter%20for%20Cilium%20Build%20and%20Releaser.svg" /></a>
[![GitHub](https://img.shields.io/github/license/layer5io/meshery-istio.svg)](LICENSE)
[![GitHub issues by-label](https://img.shields.io/github/issues/layer5io/meshery-cilium/help%20wanted.svg)](https://github.com/meshery/meshery-cilium/issues?q=is%3Aopen+is%3Aissue+label%3A"help+wanted")
[![Website](https://img.shields.io/website/https/layer5.io/meshery.svg)](https://meshery.io)
[![Twitter Follow](https://img.shields.io/twitter/follow/layer5.svg?label=Follow&style=social)](https://twitter.com/intent/follow?screen_name=mesheryio)
[![Discuss Users](https://img.shields.io/discourse/users?server=https%3A%2F%2Fdiscuss.layer5.io)](https://discuss.layer5.io)
[![Slack](https://img.shields.io/badge/Slack-@layer5.svg?logo=slack)](http://slack.layer5.io)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/3564/badge)](https://bestpractices.coreinfrastructure.org/projects/3564)

</div>

# Meshery Adapter for Cilium Service Mesh

<h2><a href="https://docs.meshery.io/">Cilium Service Mesh</h2>
<a href="https://cilium.io/">
 <img src="https://raw.githubusercontent.com/meshery/meshery-cilium/master/.github/readme/images/cilium.svg" style="margin:10px;" width="125px" 
   alt="Cilium logo" align="left" />
</a>
<p>
<a href="https://cilium.io/">Cilium</a> is open source software for providing, securing, and observing network connectivity between container workloads - cloud-native, and fueled by eBPF as Linux kernel technology. Cilium is bringing eBPF strengths to the world of service mesh. Cilium Service Mesh features eBPF-powered connectivity, traffic management, security, and observability.
 <br /> <br /> 
</p>

<p style="clear:both;">
<h2><a href="https://layer5.io/meshery">Meshery</a></h2>
<a href="https://meshery.io"><img src="https://raw.githubusercontent.com/meshery/meshery-cilium/master/.github/readme/images/meshery/meshery-logo-light.svg"
style="margin:10px;" width="125px" 
alt="Meshery - the Service Mesh Management Plane" align="left" /></a>
<a href="https://meshery.io">Meshery</a> is the multi-service mesh management plane offering lifecycle management of more types of service meshes than any other tool available today. Meshery facilitates adopting, configuring, operating and managing performance of different service meshes and incorporates the collection and display of metrics from applications running on top of any service mesh. 
<br /><br /><p align="center"><i>If you’re using Meshery or if you like the project, please <a href="https://github.com/layer5io/meshery/stargazers">★</a> star this repository to show your support! 🤩</i></p>
</p>
<br /><br />

## Join the Community!

<a name="contributing"></a><a name="community"></a>
Our projects are community-built and welcome collaboration. 👍 Be sure to see the <a href="https://layer5.io/community/handbook/repository-overview">Layer5 Community Welcome Guide</a> for a tour of resources available to you and jump into our <a href="http://slack.layer5.io">Slack</a>!

<p style="clear:both;">
<a href ="https://layer5.io/community/meshmates"><img alt="MeshMates" src=".github\readme\images\Layer5-MeshMentors-1.png" style="margin-right:10px; margin-bottom:15px;" width="28%" align="left"/></a>
<h3>Find your MeshMate</h3>

<p>MeshMates are experienced Layer5 community members, who will help you learn your way around, discover live projects and expand your community network. 
Become a <b>Meshtee</b> today!</p>

Find out more on the <a href="https://layer5.io/community">Layer5 community</a>. <br />
<br /><br />

</p>

<div>&nbsp;</div>


<a href="https://slack.meshery.io">

<picture align="right">
  <source media="(prefers-color-scheme: dark)" srcset=".github\readme\images\slack-dark-128.png"  width="110px" align="right" style="margin-left:10px;margin-top:10px;">
  <source media="(prefers-color-scheme: light)" srcset=".github\readme\images\slack-128.png" width="110px" align="right" style="margin-left:10px;padding-top:5px;">
  <img alt="Shows an illustrated light mode meshery logo in light color mode and a dark mode meshery logo dark color mode." src=".github\readme\images\slack-128.png" width="110px" align="right" style="margin-left:10px;padding-top:13px;">
</picture>
</a>

<a href="https://meshery.io/community"><img alt="Layer5 Cloud Native Community" src="https://raw.githubusercontent.com/meshery/meshery-cilium/master/.github/readme/images//community.svg" style="margin-right:8px;padding-top:5px;" width="140px" align="left" /></a>

<p>
✔️ <em><strong>Join</strong></em> any or all of the weekly meetings on <a href="https://calendar.google.com/calendar/b/1?cid=bGF5ZXI1LmlvX2VoMmFhOWRwZjFnNDBlbHZvYzc2MmpucGhzQGdyb3VwLmNhbGVuZGFyLmdvb2dsZS5jb20">community calendar</a>.<br />
✔️ <em><strong>Watch</strong></em> community <a href="https://www.youtube.com/playlist?list=PL3A-A6hPO2IMPPqVjuzgqNU5xwnFFn3n0">meeting recordings</a>.<br />
✔️ <em><strong>Access</strong></em> the <a href="https://drive.google.com/drive/u/4/folders/0ABH8aabN4WAKUk9PVA">Community Drive</a> by completing a community <a href="https://layer5.io/newcomer">Member Form</a>.<br />
✔️ <em><strong>Discuss</strong></em> in the <a href="https://discuss.layer5.io">Community Forum</a>.<br />

</p>
<p align="center">
<i>Not sure where to start?</i> Grab an open issue with the <a href="https://github.com/issues?q=is%3Aopen+is%3Aissue+archived%3Afalse+org%3Alayer5io+org%3Ameshery+org%3Aservice-mesh-performance+org%3Aservice-mesh-patterns+label%3A%22help+wanted%22+">help-wanted label</a>.</p>

## About Layer5

**Community First**

<p>The <a href="https://layer5.io">Layer5</a> community represents the largest collection of service mesh projects and their maintainers in the world.</p>

**Open Source First**

<p>We build projects to provide learning environments, deployment and operational best practices, performance benchmarks, create documentation, share networking opportunities, and more. Our shared commitment to the open source spirit pushes Layer5 projects forward.</p>

**License**

This repository and site are available as open source under the terms of the [Apache 2.0 License](https://opensource.org/licenses/Apache-2.0).
